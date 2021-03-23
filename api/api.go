package api

import "github.com/gin-gonic/gin"

func InitApiRouter(engine *gin.Engine) {
	engine.POST("/register", regi)
}
