package Proxy

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
	return &Proxy{
		Option: option,
	}
}

func (p *Proxy) Start() {
	var localConn net.Conn
	var remoteConn net.Conn
	var errchan = make(chan struct{}, 1)
	var err error
	for {
		localConn, err = net.Dial("tcp", p.Option.Local)
		if err != nil {
			log.Println("tcp1", err)
		}
		remoteConn, err = net.Dial("tcp", p.Option.Remote)
		if err != nil {
			log.Println("tcp2", err)
		}

		go func() {
			_, err := io.Copy(localConn, remoteConn)
			if err != nil {
				fmt.Println("tcp3", err)
				errchan <- struct{}{}
			}
			defer localConn.Close()
		}()
		go func() {
			_, err := io.Copy(remoteConn, localConn)
			if err != nil {
				fmt.Println("tcp4", err)
				errchan <- struct{}{}
			}
			defer remoteConn.Close()
		}()
		<-errchan
	}

}
