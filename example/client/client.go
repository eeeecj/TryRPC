package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	hello "github.com/TryRpc/api/proto/go"
	"github.com/TryRpc/pkg/Limiter"
	"google.golang.org/grpc"
	"log"
	"net"
	"sync"
	"syscall"
	"time"
)

var ClientChan = make(chan *hello.HelloRequest, 2)

func main() {
	closing := make(chan struct{})
	grpcconn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
	}
	defer grpcconn.Close()
	grpcClient := hello.NewHelloClient(grpcconn)
	var wg = &sync.WaitGroup{}
	wg.Add(2)
	go getConnect(wg, closing)
	go Solve(wg, grpcClient, closing)
	wg.Wait()
}

func getConnect(wg *sync.WaitGroup, c chan struct{}) {
	defer wg.Done()
	conn, err := net.Dial("tcp", ":8090")
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()
	//CONNECT:
	count := 0
	for {
		count++
		s, err := bufio.NewReader(conn).ReadBytes('\n')
		fmt.Println(count)
		switch err {
		case syscall.EAGAIN:
			continue
		case nil:
			var data Limiter.GrpcData
			if err := json.Unmarshal(s, &data); err != nil {
				log.Println("here", err)
			}
			r := hello.HelloRequest{Input: data.Data}
			ClientChan <- &r
			conn.Write([]byte("ok"))
		default:
			//conn, err = net.Dial("tcp", ":8090")
			//if err != nil {
			close(c)
			//	break CONNECT
			//}
			//goto CONNECT
			return
		}
	}

}

func Solve(wg *sync.WaitGroup, grpcClient hello.HelloClient, close chan struct{}) {
	defer wg.Done()
	for {
		select {
		case r := <-ClientChan:
			res, _ := grpcClient.Hello(context.Background(), r)
			fmt.Println(res)
			time.Sleep(time.Second)
		case <-close:
			return
		}
	}
}
