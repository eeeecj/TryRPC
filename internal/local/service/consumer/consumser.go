package consumer

import (
	"context"
	"fmt"
	hello "github.com/TryRpc/api/proto/go"
	"github.com/TryRpc/internal/pkg/middleware"
	"log"
	"runtime"
	"sync"
	"time"
)

type Consumer struct {
	C           int
	RequestChan chan *hello.HelloRequest
	errchan     chan struct{}
	Client      hello.HelloClient
	Limiter     *middleware.Limiter
}

func (w *Consumer) Close() {
	w.errchan <- struct{}{}
}

func (w *Consumer) Start() {
	var wg = &sync.WaitGroup{}
	wg.Add(w.C + 1)
	grpcConn, err := GetGRPC(nil)
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

func (w *Consumer) getMission(wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-w.errchan:
			return
		default:
			data := w.Limiter.ReleaseLimiter()
			r := hello.HelloRequest{Input: data.Data}
			w.RequestChan <- &r
		}
	}
}

func (w *Consumer) Solve(wg *sync.WaitGroup) {
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
