package server

import (
	"context"
	"github.com/TryRpc/component/pkg/cuszap"
	"github.com/TryRpc/internal/pkg/middleware"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"net/http"
	"time"
)

type GenericServer struct {
	*gin.Engine
	Middlewares     []string
	InsecureServing *GenericConfig
	ShutdownTimeOut time.Duration
	insecureServer  *http.Server
}

func initGenericServer(s *GenericServer) {
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		cuszap.Infof("%-6s %-s --> %s (%d handlers)", httpMethod, absolutePath, handlerName, nuHandlers)
	}
	pprof.Register(s.Engine)
	for _, m := range s.Middlewares {
		mw, ok := middleware.MiddleWares[m]
		if !ok {
			cuszap.Warnf("can not find middleware: %s", m)
			continue
		}
		cuszap.Infof("install middleware: %s", m)
		s.Use(mw)
	}
}

func (s *GenericServer) Run() error {
	s.insecureServer = &http.Server{
		Addr:    s.InsecureServing.Address(),
		Handler: s,
	}
	var eg errgroup.Group

	eg.Go(func() error {
		cuszap.Infof("Start to listening the incoming requests on http address: %s", s.InsecureServing.Address())
		if err := s.insecureServer.ListenAndServe(); err != nil {
			cuszap.Fatalf(err.Error())
			return err
		}
		cuszap.Infof("Server on %s stopped", s.InsecureServing.Address())
		return nil
	})

	if err := eg.Wait(); err != nil {
		cuszap.Fatal(err.Error())
	}
	return nil
}

func (s *GenericServer) Close() {
	ctx, cancel := context.WithTimeout(context.Background(), s.ShutdownTimeOut*time.Second)
	defer cancel()
	if err := s.insecureServer.Shutdown(ctx); err != nil {
		cuszap.Warnf("Shutdown insecure server failed: %s", err.Error())
	}
}
