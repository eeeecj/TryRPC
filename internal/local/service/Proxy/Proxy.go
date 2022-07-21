package Proxy

import (
	"bufio"
	"fmt"
	"github.com/TryRpc/pkg/service"
	"io"
	"log"
	"time"
)

type Proxy struct {
	closing chan struct{}
}

func NewProxy() *Proxy {
	return &Proxy{closing: make(chan struct{}, 1)}
}

const NewConnection = "NewConnection\n"

var option = NewProxyOption()

func (p *Proxy) Start() {
	controllerConnection(p.closing)
}

func (p *Proxy) Close() {
	localpool.Close()
	remotePool.Close()
	p.closing <- struct{}{}
	time.Sleep(time.Second * 3)
}
func controllerConnection(closing chan struct{}) {
	controllerConn, err := service.CreateConn(option.Controller)
	if err != nil {
		log.Println("连接控制器失败：", err)
		return
	}
	defer controllerConn.Close()
	log.Printf("连接控制器成功:%s", option.Controller)
	for {
		select {
		case <-closing:
			fmt.Println("closing")
			return
		default:
			buf, err := bufio.NewReader(controllerConn).ReadString('\n')
			if err != nil {
				log.Println(err)
				return
			}
			if buf == NewConnection {
				go proxyConn(closing)
			}
		}
	}
}

func proxyConn(closing chan struct{}) {
	remoteConn, err := remotePool.Get()
	if err != nil {
		panic(err)
	}
	log.Println("连接远程代理成功:", option.Remote)
	localConn, err := localpool.Get()
	if err != nil {
		panic(err)
	}
	log.Println("连接本地服务成功：", option.Local)
	go func() {
		io.Copy(localConn, remoteConn)
		defer localConn.Close()
	}()

	go func() {
		io.Copy(remoteConn, localConn)
		defer remoteConn.Close()
	}()
}
