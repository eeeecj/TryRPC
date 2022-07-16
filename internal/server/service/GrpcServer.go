package service

import (
	"encoding/json"
	"fmt"
	"github.com/TryRpc/internal/server/middlewares"
	"log"
	"net"
)

type Proxy struct {
	c        int
	sendChan chan struct{}
}

func NewProxy() *Proxy {
	p := &Proxy{c: 2, sendChan: make(chan struct{}, 2)}
	for i := 0; i < 2; i++ {
		p.sendChan <- struct{}{}
	}
	return p
}

const bufsize = 1024

func (p *Proxy) Run() {
	lis, err := net.Listen("tcp", ":8090")
	fmt.Println("sss")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := lis.Accept()
		fmt.Println("sss")
		if err != nil {
			log.Println(err)
		}
		go p.handler(conn)
	}
}

func (p *Proxy) handler(conn net.Conn) {
	if conn == nil {
		log.Println("nil conn pass")
	}
	go p.sendMsg(conn)
	go p.rcvMsg(conn)
}

func (p *Proxy) sendMsg(conn net.Conn) {
	for {
		<-p.sendChan
		c := middlewares.DefaultLimiter.ReleaseLimiter()
		res, err := json.Marshal(c)
		if err != nil {
			log.Println("json :", err)
		}
		_, err = conn.Write(res)
		_, err = conn.Write([]byte{'\n'})
		if err != nil {
			log.Println(err)
			conn.Close()
			return
		}
	}

}

func (p *Proxy) rcvMsg(conn net.Conn) {
	for {
		var b [2]byte
		_, err := conn.Read(b[:])
		if err != nil {
			log.Println(err)
			conn.Close()
			return
		}
		if string(b[:]) == "ok" {
			p.sendChan <- struct{}{}
		}
	}
}
