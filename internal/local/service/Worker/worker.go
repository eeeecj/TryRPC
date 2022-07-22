package Worker

import (
	"context"
	"fmt"
	hello "github.com/TryRpc/api/proto/go"
	"github.com/TryRpc/pkg/Limiter"
	"google.golang.org/grpc"
	"log"
	"runtime"
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

func (w *Worker) Close() {
	w.errchan <- struct{}{}
}

func (w *Worker) Start() {
	var wg = &sync.WaitGroup{}
	wg.Add(w.C + 1)
	grpcConn, err := grpc.Dial(":1234", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
	}
	client := hello.NewHelloClient(grpcConn)
	w.Client = client
	go w.getMission(wg)
	for i := 0; i < w.C; i++ {
		go w.Solve(wg)
	}
	wg.Wait()
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
		select {
		case <-w.errchan:
			runtime.Goexit()
		case r := <-w.RequestChan:
			res, err := w.Client.Hello(context.Background(), r)
			time.Sleep(3 * time.Second)
			if err != nil {
				log.Println(err)
				w.errchan <- struct{}{}
				return
			}
			fmt.Println(res)
		}

	}
}
