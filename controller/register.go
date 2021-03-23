package controller

import "github.com/gin-gonic/gin"

func Register(c *gin.Context) {
	username, password := c.PostForm("username"), c.PostForm("password")

}
