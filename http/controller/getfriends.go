package controller

import (
	"github.com/gin-gonic/gin"
	"lim/http/service"
	"strconv"
)

func GetFriends(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(401, gin.H{
			"err": err,
		})
		return
	}

	friendIds := service.GetFriends(id)
	var friendNames []string
	for _, x := range friendIds {
		user, _ := service.GetUser(x)
		friendNames = append(friendNames, user.Username)
	}
	c.JSON(200, gin.H{
		"friendIds":   friendIds,
		"friendNames": friendNames,
	})
}
