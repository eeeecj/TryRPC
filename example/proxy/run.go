package main

import (
	"io"
	"net"
)

type Server struct {
	Addr   string
	Target string
}

func main() {
	s := &Server{
		Addr:   ":8777",
		Target: ":20012",
	}

	ser, err := net.Listen("tcp", s.Addr)
	if err != nil {
		println(err)
	}

	tar, err := net.Listen("tcp", s.Target)
	if err != nil {
		println(err)
	}
	for {
		serconn, _ := ser.Accept()
		go func() {
			tarconn, _ := tar.Accept()

			go func() {
				io.Copy(tarconn, serconn)
				defer tarconn.Close()
			}()
			go func() {
				io.Copy(serconn, tarconn)
				defer serconn.Close()
			}()
		}()
	}
}
