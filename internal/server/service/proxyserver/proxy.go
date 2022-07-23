package proxyserver

import (
	"context"
	"fmt"
	"github.com/TryRpc/component/pkg/cuszap"
	"github.com/TryRpc/pkg/utils"
	"io"
	"log"
	"net"
	"sync"
	"time"
)

type ProxyServer struct {
	RemoteServing     *RemoteServing
	LocalServing      *LocalServing
	ControllerServing *ControllerServing
}

const KeepAlive = "KeepAlive\n"
const NewConnection = "NewConnection\n"

var ctx, cancel = context.WithCancel(context.Background())

func (p *ProxyServer) Start() {
	var wg = &sync.WaitGroup{}
	wg.Add(1)
	go p.controlListen(wg)
	go p.localListen(wg)
	go p.proxyListen(wg)
	wg.Wait()
	fmt.Println("this is wg")
}

func (p *ProxyServer) Close() {
	cancel()
}

var controllerConn net.Conn
var localConn net.Conn

func (p *ProxyServer) controlListen(wg *sync.WaitGroup) {
	defer wg.Done()
	lis, err := utils.CreateListenWithTLS(p.ControllerServing.Address(),
		p.ControllerServing.CertKey.CertFile, p.ControllerServing.CertKey.KeyFile, p.ControllerServing.CertKey.CaFile)
	if err != nil {
		cuszap.Fatalf("controller listen fail:%v", err)
	}
	defer lis.Close()
	cuszap.Infof("控制中心监听中：http://localhost:%s\n", p.ControllerServing.Address())
	for {
		select {
		case <-ctx.Done():
			fmt.Println("timeout")
			return
		default:
			controllerConn, err = lis.Accept()
			if err != nil {
				log.Println("控制中心请求失败：", err)
				return
			}
			defer controllerConn.Close()
			go keepAlive(controllerConn)
		}
	}
}

func (p *ProxyServer) localListen(wg *sync.WaitGroup) {
	defer wg.Done()
	lis, err := utils.CreateListen(p.LocalServing.Address())
	if err != nil {
		log.Println("local listen fail:", err)
		return
	}
	cuszap.Infof("服务器代理端口成功：http://localhost:%s", p.LocalServing.Address())
	for {
		select {
		case <-ctx.Done():
			fmt.Println("timeout")
			return
		default:
			localConn, err = lis.Accept()
			if err != nil {
				log.Println("本地监听失败：", err)
				return
			}
			_, err := controllerConn.Write([]byte(NewConnection))
			if err != nil {
				log.Println("创建新连接失败：", err)
			}
		}
	}
}

func (p *ProxyServer) proxyListen(wg *sync.WaitGroup) {
	defer wg.Done()
	proxyLis, err := utils.CreateListenWithTLS(p.RemoteServing.Address(),
		p.RemoteServing.CertKey.CertFile, p.RemoteServing.CertKey.KeyFile, p.ControllerServing.CertKey.CaFile)
	if err != nil {
		cuszap.Infof("代理服务启动失败：%v", err)
		return
	}
	cuszap.Infof("代理服务启动成功:http://localhost:%s", p.RemoteServing.Address())
	for {
		select {
		case <-ctx.Done():
			fmt.Println("timeout")
			return
		default:
			proxyConn, err := proxyLis.Accept()
			if err != nil {
				log.Println("代理接收请求失败：", err)
				return
			}
			go io.Copy(localConn, proxyConn)
			go io.Copy(proxyConn, localConn)
		}
	}
}

func keepAlive(conn net.Conn) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			_, err := conn.Write([]byte(KeepAlive))
			if err != nil {
				log.Println("keep alive error:", err)
				return
			}
			time.Sleep(3 * time.Second)
		}
	}
}
