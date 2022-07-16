package service

import (
	"fmt"
	"io"
	"log"
	"net"
)

type Proxy struct {
	Option *ProxyOption
}

func NewProxy(option *ProxyOption) *Proxy {
	return &Proxy{Option: option}
}

func (p *Proxy) startServer(Addr string) net.Listener {
	lis, err := net.Listen("tcp", Addr)
	if err != nil {
		log.Println("start server:", err)
	}
	return lis
}
func (p *Proxy) Start() {
	localLis := p.startServer(p.Option.Local)
	remoteLis := p.startServer(p.Option.Remote)
	for {
		fmt.Println("sajslkalksa")
		remoteConn, err := remoteLis.Accept()
		if err != nil {
			log.Println("server1", err)
		}
		go func() {
			localConn, err := localLis.Accept()
			if err != nil {
				log.Println("server2", err)
			}
			go func() {
				io.Copy(localConn, remoteConn)
				defer localConn.Close()
			}()
			go func() {
				io.Copy(remoteConn, localConn)
				defer remoteConn.Close()
			}()
		}()
	}
}
