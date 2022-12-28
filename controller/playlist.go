package controller

import (
	"github.com/96368a/NewApi/services"
	"github.com/96368a/NewApi/utils"
	"github.com/gin-gonic/gin"
)

func GetPlaylistToplist(c *gin.Context) {
	toplist := services.GetPlaylistToplist()
	utils.Success(c, gin.H{
		"list": toplist,
	})

}
