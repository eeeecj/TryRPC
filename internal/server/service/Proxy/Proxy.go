package Proxy

import (
	"context"
	"fmt"
	"github.com/TryRpc/pkg/service"
	"io"
	"log"
	"net"
	"sync"
	"time"
)

var option = NewProxyOption()

const KeepAlive = "KeepAlive\n"
const NewConnection = "NewConnection\n"

var ctx, cancel = context.WithCancel(context.Background())

type Proxy struct {
}

func Axxx() {
	for {
		select {
		case <-time.After(time.Second * 5):
			fmt.Println("timeout")
			return
		case <-ctx.Done():
			fmt.Println("ctx done")
			return
		}
	}
}

func NewProxy() *Proxy {
	return &Proxy{}
}
func (p *Proxy) Start() {
	var wg = &sync.WaitGroup{}
	wg.Add(1)
	go controlListen(wg)
	go localListen(wg)
	go proxyListen(wg)
	wg.Wait()
	fmt.Println("this is wg")
}

func (p *Proxy) Close() {
	cancel()
}

var controllerConn *net.TCPConn
var localConn *net.TCPConn

func controlListen(wg *sync.WaitGroup) {
	defer wg.Done()
	lis, err := service.CreateListen(option.Controller)
	if err != nil {
		log.Fatalln("controller listen fail:", err)
	}
	defer lis.Close()
	log.Printf("控制中心监听中：http://localhost:%s\n", option.Controller)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("timeout")
			return
		default:
			controllerConn, err = lis.AcceptTCP()
			if err != nil {
				log.Println("控制中心请求失败：", err)
				return
			}
			defer controllerConn.Close()
			go keepAlive(controllerConn)
		}
	}
}

func localListen(wg *sync.WaitGroup) {
	defer wg.Done()
	lis, err := service.CreateListen(option.Local)
	if err != nil {
		log.Println("local listen fail:", err)
		return
	}
	log.Printf("服务器代理端口成功：http://localhost:%s", option.Local)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("timeout")
			return
		default:
			localConn, err = lis.AcceptTCP()
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

func proxyListen(wg *sync.WaitGroup) {
	defer wg.Done()
	proxyLis, err := service.CreateListen(option.Remote)
	if err != nil {
		log.Println("代理服务启动失败：", err)
		return
	}
	log.Printf("代理服务启动成功:http://localhost:%s", option.Remote)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("timeout")
			return
		default:
			proxyConn, err := proxyLis.AcceptTCP()
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
