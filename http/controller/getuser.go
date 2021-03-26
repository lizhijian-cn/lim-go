package controller

import (
	"github.com/gin-gonic/gin"
	"lim/http/service"
	"strconv"
)

func GetUser(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(401, gin.H{
			"err": "invalid user id",
		})
		return
	}
	user, err := service.GetUser(id)
	if err != nil {
		c.JSON(401, gin.H{
			"err": err,
		})
		return
	}
	c.JSON(200, gin.H{
		"msg":      "successfully get user",
		"username": user.Username,
	})
}
