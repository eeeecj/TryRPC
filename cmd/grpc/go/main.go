package main

import (
	"context"
	"fmt"
	hello "github.com/TryRpc/api/proto/go"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Printf("failed to listen:%s", err.Error())
	}
	server := grpc.NewServer()
	hello.RegisterHelloServer(server, new(GRPCFactory))
	server.Serve(lis)
}

type GRPCFactory struct {
}

func (g *GRPCFactory) Hello(ctx context.Context, request *hello.HelloRequest) (*hello.HelloResponse, error) {
	fmt.Println(request)
	reply := &hello.HelloResponse{Output: "hello " + request.Input + " Test"}
	return reply, nil
}
