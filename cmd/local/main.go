package main

import (
	"context"
	"encoding/json"
	"fmt"
	hello "github.com/TryRpc/api/proto"
	"github.com/TryRpc/internal/server/middlewares"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

var clientChan = make(chan *hello.HelloRequest, 2)

func main() {
	client := &http.Client{}
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()
	grpcClient := hello.NewHelloClient(conn)
	var wg = &sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			req, _ := http.NewRequest("POST", "http://127.0.0.1:8080/v1/sims/get", nil)
			fmt.Println("sss")
			response, err := client.Do(req)
			if err != nil {
				log.Println("threer", err)
				time.Sleep(time.Second)
				continue
			}
			if response.StatusCode == 200 {
				body, _ := ioutil.ReadAll(response.Body)
				var data middlewares.GrpcData
				if err := json.Unmarshal(body, &data); err != nil {
					log.Println("here", err)
				}
				r := hello.HelloRequest{Input: data.Data}
				fmt.Println(r)
				clientChan <- &r
			}
			time.Sleep(time.Second)
		}
	}()
	go Solve(wg, grpcClient)
	wg.Wait()
}

func Solve(wg *sync.WaitGroup, grpcClient hello.HelloClient) {
	defer wg.Done()
	for {
		r := <-clientChan
		fmt.Println(r)
		res, _ := grpcClient.Hello(context.Background(), r)
		fmt.Println(res)

	}
}
