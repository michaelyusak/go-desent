package server

import (
	"michaelyusak/go-desent.git/handler"

	"github.com/gin-gonic/gin"
)

type routerOpts struct {
	handler struct {
		common *handler.Common
	}
}

func newRouter() *gin.Engine {
	return createRouter(routerOpts{
		handler: struct {
			common *handler.Common
		}{
			common: &handler.Common{},
		},
	},
		[]string{},
	)
}

func createRouter(opts routerOpts, allowedOrigins []string) *gin.Engine {
	router := gin.New()

	router.Use(
		gin.Recovery(),
	)

	commonRouting(router, opts.handler.common)

	return router
}

func commonRouting(router *gin.Engine, handler *handler.Common) {
	router.GET("/ping", handler.Ping)
	router.POST("/echo", handler.Echo)
}
