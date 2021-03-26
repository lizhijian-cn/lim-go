package controller

import (
	"github.com/gin-gonic/gin"
	"lim/http/service"
	"lim/util"
	"log"
	"strconv"
)

type LoginForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	var form LoginForm
	if err := c.BindJSON(&form); err != nil {
		c.JSON(401, gin.H{
			"err": err.Error(),
		})
		return
	}
	user, err := service.Login(form.Username, form.Password)
	if err != nil {
		c.JSON(401, gin.H{
			"err": err.Error(),
		})
		return
	}

	_, err = util.GenerateToken(user.Username, 1<<20)
	if err != nil {
		log.Fatalf("failed to generate token: %s", err)
		c.JSON(401, gin.H{
			"err": "other errors",
		})
		return
	}
	c.JSON(200, gin.H{
		//"token": token,
		"msg":    "successfully login",
		"userId": strconv.FormatInt(user.Id, 10),
	})
}
