package main

import (
	"bufio"
	"io"
	"log"
	"net"
)

type Server struct {
	Addr   string
	Target string
}

func (s *Server) Serve(l net.Listener) {
	for {
		conn, err := l.Accept()
		if err != nil {
			panic(err)
		} else {
			go s.ServeConn(conn)
		}
	}
}

func (s *Server) ListenAndServe() error {
	l, err := net.Listen("tcp", s.Addr)
	if err != nil {
		return err
	}
	s.Serve(l)
	target, err := net.Listen("tcp", s.Target)
	if err != nil {
		return err
	}
	s.Serve(target)
	return nil
}

func (s *Server) ServeConn(conn net.Conn) {
	rbuf := bufio.NewReader(conn)
	wbuf := bufio.NewWriter(conn)
	s.handleConnect(wbuf, rbuf)
}
func (s *Server) handleConnect(w *bufio.Writer, r *bufio.Reader) {

	if err != nil {
		log.Println(err)
	}
	go s.proxy(target, r)
	go s.proxy(w, target)
}
func (s *Server) proxy(dest io.Writer, src io.Reader) {
	n, err := io.Copy(dest, src)
	if err != nil {
		log.Println(n, err)
	}
}
