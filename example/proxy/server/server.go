package main

import (
	"fmt"
	"io/ioutil"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":7809")
	if err != nil {
		panic(err)
	}
	for {
		conn, _ := lis.Accept()
		fmt.Println("a connect")
		buf, _ := ioutil.ReadAll(conn)
		fmt.Println(string(buf))
		conn.Write([]byte("hello world"))
	}
}
