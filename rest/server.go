package rest

import (
	"frame/pkg/grace"
	"frame/pkg/middleware"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"time"
)

type HttpServer struct {
	addr string
}

func NewHttpServer(addr string) *HttpServer {
	return &HttpServer{
		addr: addr,
	}
}

// 启动服务
func (server *HttpServer) Start() {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = io.Discard

	r := gin.Default()
	r.Use(middleware.SetCROSOptions)

	v1 := r.Group("/v1")
	{
		v1.GET("/ping", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "ok")
		})
	}

	// 优雅的重启
	grace.HttpRun(server.addr, r, 10*time.Minute)
}
