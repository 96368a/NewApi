package controller

import (
	"github.com/96368a/NewApi/services"
	"github.com/96368a/NewApi/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetArtistAlbum(c *gin.Context) {
	idStr := c.Query("id")
	limitStr := c.DefaultQuery("limit", "30")
	if idStr == "" {
		utils.Fail(c, 400, "参数错误")
		return
	}
	id, _ := strconv.Atoi(idStr)
	if id <= 0 {
		utils.Fail(c, 400, "参数错误")
		return
	}
	limit, _ := strconv.Atoi(limitStr)
	if limit <= 0 {
		limit = 30
	}
	albums := services.GetArtistAlbum(int64(id), limit)
	artist := services.GetOneArtist(int64(id))
	if albums == nil {
		utils.Fail(c, 400, "参数错误")
		return
	}
	utils.Success(c, gin.H{
		"artist": artist,
		"hotAlbums": gin.H{
			"albums": albums,
		},
	})

}
