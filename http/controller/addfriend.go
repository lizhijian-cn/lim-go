package controller

import (
	"github.com/gin-gonic/gin"
	"lim/http/service"
	"strconv"
)

type addFriendForm struct {
	Aid string `json:"aid"`
	Bid string `json:"bid"`
}

func AddFriend(c *gin.Context) {
	var form addFriendForm
	if err := c.BindJSON(&form); err != nil {
		c.JSON(401, gin.H{
			"err": err.Error(),
		})
		return
	}
	aid, err := strconv.ParseInt(form.Aid, 10, 64)
	if err != nil {
		c.JSON(401, gin.H{
			"err": err.Error(),
		})
		return
	}
	bid, err := strconv.ParseInt(form.Bid, 10, 64)
	if err != nil {
		c.JSON(401, gin.H{
			"err": err.Error(),
		})
		return
	}

	err = service.AddFriend(aid, bid)
	if err != nil {
		c.JSON(401, gin.H{
			"err": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"msg": "successfully add friend",
	})
}
