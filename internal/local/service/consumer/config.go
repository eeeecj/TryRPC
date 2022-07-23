package consumer

import (
	"fmt"
	hello "github.com/TryRpc/api/proto/go"
	"github.com/TryRpc/internal/local/consumer"
	"github.com/TryRpc/internal/pkg/middleware"
	"net"
	"strconv"
)

type ConsumerConfig struct {
	MaxConsumer int
	BindAddr    string
	BindPort    int
}

func NewConsumerConfig() *ConsumerConfig {
	return &ConsumerConfig{MaxConsumer: 2}
}

func (g *ConsumerConfig) ToApply(s *consumer.ConsumerOption) error {
	g.MaxConsumer = s.MaxConsumer
	g.BindAddr = s.BindAddr
	g.BindPort = s.BindPort
	return nil
}
func (g *ConsumerConfig) Address() string {
	return net.JoinHostPort(g.BindAddr, strconv.Itoa(g.BindPort))
}
func (s *ConsumerConfig) New() (*Consumer, error) {
	fmt.Printf("%+v", s)
	limiter, _ := middleware.GetLimiter(nil)
	GetGRPC(s)
	server := &Consumer{
		C:           s.MaxConsumer,
		RequestChan: make(chan *hello.HelloRequest, s.MaxConsumer),
		errchan:     make(chan struct{}, 1),
		Limiter:     limiter,
	}
	return server, nil
}
