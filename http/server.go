package httpserver

import (
	"github.com/aceld/zinx/zlog"
	"github.com/gin-gonic/gin"
)

func StartWebServer(addr string) {
	router := gin.Default()
	InitApiRouter(router)
	err := router.Run(addr)
	if err != nil {
		zlog.Fatalf("failed to start web server: %s", err)
		return
	}
}
