package proxy

import (
	"bufio"
	"fmt"
	"github.com/TryRpc/component/pkg/cuszap"
	"github.com/TryRpc/pkg/utils"
	"github.com/fatih/pool"
	"io"
	"log"
	"net"
	"time"
)

type ProxyServer struct {
	RemoteServing     *RemoteClientServing
	LocalServing      *LocalClientServing
	ControllerServing *ControllerClientServing
	closing           chan struct{}
	localPool         pool.Pool
	remotePool        pool.Pool
	PoolSize          int
}

func NewProxy() *ProxyServer {
	return &ProxyServer{closing: make(chan struct{}, 1)}
}

const NewConnection = "NewConnection\n"

func (p *ProxyServer) Start() {
	p.localPool, _ = pool.NewChannelPool(0, p.PoolSize, p.localFactory)
	p.remotePool, _ = pool.NewChannelPool(0, p.PoolSize, p.remoteFactory)
	p.controllerConnection()
}

func (p *ProxyServer) localFactory() (net.Conn, error) {
	return utils.CreateConn(p.LocalServing.Address())
}
func (p *ProxyServer) remoteFactory() (net.Conn, error) {
	return utils.CreateConnWithTLS(p.RemoteServing.Address(), p.RemoteServing.CertKey.CertFile,
		p.RemoteServing.CertKey.KeyFile, p.RemoteServing.CertKey.CaFile)
}
func (p *ProxyServer) Close() {
	p.localPool.Close()
	p.remotePool.Close()
	p.closing <- struct{}{}
	time.Sleep(time.Second * 3)
}
func (p *ProxyServer) controllerConnection() {
	controllerConn, err := utils.CreateConnWithTLS(p.ControllerServing.Address(),
		p.ControllerServing.CertKey.CertFile, p.ControllerServing.CertKey.KeyFile, p.ControllerServing.CertKey.CaFile)
	if err != nil {
		log.Println("连接控制器失败：", err)
		return
	}
	defer controllerConn.Close()
	log.Printf("连接控制器成功:%s", p.ControllerServing.Address())
	for {
		select {
		case <-p.closing:
			fmt.Println("closing")
			return
		default:
			buf, err := bufio.NewReader(controllerConn).ReadString('\n')
			if err != nil {
				cuszap.Infof("%v", err)
				return
			}
			if buf == NewConnection {
				go p.proxyConn()
			}
		}
	}
}

func (p *ProxyServer) proxyConn() {
	remoteConn, err := p.remotePool.Get()
	if err != nil {
		panic(err)
	}
	cuszap.Info("连接远程代理成功:" + p.RemoteServing.Address())
	localConn, err := p.localPool.Get()
	if err != nil {
		panic(err)
	}
	cuszap.Info("连接本地服务成功：" + p.LocalServing.Address())
	go func() {
		io.Copy(localConn, remoteConn)
		defer localConn.Close()
	}()

	go func() {
		io.Copy(remoteConn, localConn)
		defer remoteConn.Close()
	}()
}
