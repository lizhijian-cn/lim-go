package httpserver

import (
	"github.com/gin-gonic/gin"
	"lim/http/controller"
)

func InitApiRouter(engine *gin.Engine) {
	engine.POST("/register", controller.Register)
	engine.POST("/login", controller.Login)
	engine.GET("/user/:id", controller.GetUser)

	engine.GET("/relation/:id", controller.GetFriends)
	engine.POST("/relation/add", controller.AddFriend)
}
