package server

import (
	"github.com/TryRpc/internal/server/service"
	"github.com/TryRpc/pkg/app"
	"github.com/gin-gonic/gin"
)

func New() *app.App {
	return &app.App{
		gin.New(),
		[]string{},
		service.NewProxy(),
	}
}
