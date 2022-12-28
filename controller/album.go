package controller

import (
	"github.com/96368a/NewApi/services"
	"github.com/96368a/NewApi/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetAlbumSongs(c *gin.Context) {
	idStr := c.Query("id")
	if idStr == "" {
		utils.Fail(c, 400, "参数错误")
		return
	}
	id, _ := strconv.Atoi(idStr)
	if id <= 0 {
		utils.Fail(c, 400, "参数错误")
		return
	}
	album := services.GetOneAlbum(int64(id))
	songs := services.GetAlbumSongs(int64(id), 9999)
	utils.Success(c, gin.H{
		"songs": songs,
		"album": album,
	})
}

func GetNewAlbum(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "30")
	limit, _ := strconv.Atoi(limitStr)
	offsetStr := c.DefaultQuery("offset", "0")
	offset, _ := strconv.Atoi(offsetStr)
	if limit <= 0 {
		limit = 30
	}
	if offset < 0 {
		offset = 0
	}
	albums := services.GetNewAlbum(int64(limit), int64(offset))
	utils.Success(c, gin.H{
		"albums": albums,
	})

}
