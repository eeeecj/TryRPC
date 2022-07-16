package Proxy

import (
	"bufio"
	"github.com/TryRpc/pkg/service"
	"io"
	"log"
)

type Proxy struct {
}

func NewProxy() *Proxy {
	return &Proxy{}
}

const NewConnection = "NewConnection\n"

var option = NewProxyOption()

func (p *Proxy) Start() {
	controllerConnection()
}

func controllerConnection() {
	var closing = make(chan struct{}, 1)
	controllerConn, err := service.CreateConn(option.Controller)
	if err != nil {
		log.Println("连接控制器失败：", err)
		return
	}
	defer controllerConn.Close()
	log.Printf("连接控制器成功:%s", option.Controller)
	for {
		buf, err := bufio.NewReader(controllerConn).ReadString('\n')
		switch err {
		case nil:
			if buf == NewConnection {
				go proxyConn(closing)
			}
		default:
			log.Println(err)
			return
		}
	}
}

func proxyConn(closing chan struct{}) {
	remoteConn, err := service.CreateConn(option.Remote)
	if err != nil {
		panic(err)
	}
	log.Println("连接远程代理成功:", option.Remote)
	localConn, err := service.CreateConn(option.Local)
	if err != nil {
		panic(err)
	}
	log.Println("连接本地服务成功：", err)
	go io.Copy(localConn, remoteConn)

	go io.Copy(remoteConn, localConn)
}
