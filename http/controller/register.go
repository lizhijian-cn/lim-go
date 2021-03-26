package controller

import (
	"github.com/gin-gonic/gin"
	"lim/http/service"
)

func Register(c *gin.Context) {
	username, password := c.PostForm("username"), c.PostForm("password")
	err := service.Register(username, password)
	if err != nil {
		c.JSON(401, gin.H{
			"err": err,
		})
	} else {
		c.JSON(200, gin.H{
			"msg": "successfully register",
		})
	}
}
