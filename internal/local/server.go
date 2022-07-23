package local

import (
	"github.com/TryRpc/component/pkg/cuszap"
	"github.com/TryRpc/component/pkg/shutdown"
	"github.com/TryRpc/component/pkg/shutdown/posixsignal"
	"github.com/TryRpc/internal/local/config"
	"github.com/TryRpc/internal/local/service/consumer"
	"github.com/TryRpc/internal/local/service/proxy"
	"github.com/TryRpc/internal/local/service/server"
)

type apiServer struct {
	gs            *shutdown.GracefulShutdown
	genericServer *server.GenericServer
	proxyServer   *proxy.ProxyServer
	consumer      *consumer.Consumer
}

func createAPIServer(cfg *config.Config) (*apiServer, error) {
	gs := shutdown.New()
	gs.AddShutdownManager(posixsignal.NewPosixSignalManager())

	genericConfig, err := buildGenericConfig(cfg)
	if err != nil {
		return nil, err
	}
	genericServer, err := genericConfig.New()
	if err != nil {
		return nil, err
	}
	proxyConfig, err := buildProxyConfig(cfg)
	if err != nil {
		return nil, err
	}
	proxyServer, err := proxyConfig.New()
	if err != nil {
		return nil, err
	}
	consumerConfig, err := buildConsumerConfig(cfg)
	if err != nil {
		return nil, err
	}
	consumer, err := consumerConfig.New()
	server := &apiServer{
		gs:            gs,
		genericServer: genericServer,
		proxyServer:   proxyServer,
		consumer:      consumer,
	}
	return server, nil
}

func (s *apiServer) PrepareRun() *apiServer {
	initRouter(s.genericServer.Engine)
	s.gs.AddShutdownCallback(shutdown.ShutdownFunc(func(string) error {
		s.genericServer.Close()
		s.proxyServer.Close()
		s.consumer.Close()
		return nil
	}))
	return s
}

func (s *apiServer) Run() error {
	go s.consumer.Start()
	go s.proxyServer.Start()
	if err := s.gs.Start(); err != nil {
		cuszap.Fatalf("start shutdown manager failed: %s", err.Error())
	}
	return s.genericServer.Run()
}

func buildGenericConfig(cfg *config.Config) (genericConfig *server.GenericConfig, lastErr error) {
	genericConfig = server.NewGenericConfig()
	if lastErr = genericConfig.ToApply(cfg.GenericServerServing); lastErr != nil {
		return
	}
	return
}

func buildProxyConfig(cfg *config.Config) (proxyConfig *proxy.ProxyClientConfig, lastErr error) {
	proxyConfig = proxy.NewProxyClientConfig()
	if lastErr = proxyConfig.ToApply(cfg.ProxyServerServing); lastErr != nil {
		return
	}
	return
}

func buildConsumerConfig(cfg *config.Config) (consumerConfig *consumer.ConsumerConfig, lastErr error) {
	consumerConfig = consumer.NewConsumerConfig()
	if lastErr = consumerConfig.ToApply(cfg.ConsumerOptions); lastErr != nil {
		return
	}
	return
}
