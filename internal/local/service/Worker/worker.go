package Worker

import (
	"context"
	"fmt"
	hello "github.com/TryRpc/api/proto"
	"github.com/TryRpc/pkg/Limiter"
	"google.golang.org/grpc"
	"log"
	"sync"
	"time"
)

type Worker struct {
	C           int
	RequestChan chan *hello.HelloRequest
	errchan     chan struct{}
	Client      hello.HelloClient
}

func NewWorker() *Worker {
	return &Worker{
		C:           2,
		RequestChan: make(chan *hello.HelloRequest, 2),
		errchan:     make(chan struct{}, 1),
	}
}

func (w *Worker) Start() {
	for {
		var wg = &sync.WaitGroup{}
		wg.Add(2)
		grpcConn, err := grpc.Dial(":1234", grpc.WithInsecure())
		if err != nil {
			log.Println(err)
		}
		client := hello.NewHelloClient(grpcConn)
		w.Client = client
		go w.getMission(wg)
		go w.Solve(wg)
		wg.Wait()
	}
}

func (w *Worker) getMission(wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-w.errchan:
			return
		default:
			data := Limiter.DefaultLimiter.ReleaseLimiter()
			r := hello.HelloRequest{Input: data.Data}
			w.RequestChan <- &r
		}
	}
}

func (w *Worker) Solve(wg *sync.WaitGroup) {
	defer wg.Done()
	for {

		r := <-w.RequestChan
		res, err := w.Client.Hello(context.Background(), r)
		if err != nil {
			log.Println(err)
			w.errchan <- struct{}{}
		}
		fmt.Println(res)
		time.Sleep(time.Second)

	}
}
