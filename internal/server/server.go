package server

import (
	"github.com/TryRpc/component/pkg/cuszap"
	"github.com/TryRpc/component/pkg/shutdown"
	"github.com/TryRpc/component/pkg/shutdown/posixsignal"
	"github.com/TryRpc/internal/server/config"
	"github.com/TryRpc/internal/server/service/genericserver"
	"github.com/TryRpc/internal/server/service/proxyserver"
)

type apiServer struct {
	gs            *shutdown.GracefulShutdown
	genericServer *genericserver.GenericServer
	proxyServer   *proxyserver.ProxyServer
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
	server := &apiServer{
		gs:            gs,
		genericServer: genericServer,
		proxyServer:   proxyServer,
	}
	return server, nil
}

func (s *apiServer) PrepareRun() *apiServer {
	initRouter(s.genericServer.Engine, s.genericServer.LocalServing.Address())
	s.gs.AddShutdownCallback(shutdown.ShutdownFunc(func(string) error {
		s.genericServer.Close()
		s.proxyServer.Close()
		return nil
	}))
	return s
}

func (s *apiServer) Run() error {
	go s.proxyServer.Start()

	if err := s.gs.Start(); err != nil {
		cuszap.Fatalf("start shutdown manager failed: %s", err.Error())
	}
	return s.genericServer.Run()
}

func buildGenericConfig(cfg *config.Config) (genericConfig *genericserver.GenericConfig, lastErr error) {
	genericConfig = genericserver.NewGenericConfig()
	if lastErr = genericConfig.ToApply(cfg.GenericServerRunOptions); lastErr != nil {
		return
	}
	if lastErr = genericConfig.LocalToApply(cfg.ProxyServer.LocalServer); lastErr != nil {
		return
	}
	return
}

func buildProxyConfig(cfg *config.Config) (proxyConfig *proxyserver.ProxyConfig, lastErr error) {
	proxyConfig = proxyserver.NewProxyConfig()
	if lastErr = proxyConfig.ToApply(cfg.ProxyServer); lastErr != nil {
		return
	}
	return
}
