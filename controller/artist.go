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

func GetArtist(c *gin.Context) {
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
	artist := services.GetOneArtist(int64(id))
	songs := services.GetArtistHotSongs(int64(id))
	if artist == nil {
		utils.Fail(c, 400, "参数错误")
		return
	}
	utils.Success(c, gin.H{
		"artist":   artist,
		"hotSongs": songs,
	})
}

func GetArtistSongs(c *gin.Context) {
	idStr := c.Query("id")
	limitStr := c.DefaultQuery("limit", "30")
	offsetStr := c.DefaultQuery("offset", "0")
	if idStr == "" {
		utils.Fail(c, 400, "参数错误")
		return
	}
	limit, _ := strconv.Atoi(limitStr)
	offset, _ := strconv.Atoi(offsetStr)
	id, _ := strconv.Atoi(idStr)
	if id <= 0 {
		utils.Fail(c, 400, "参数错误")
		return
	}
	songs := services.GetArtistSongs(int64(id), int64(limit), int64(offset))
	if songs == nil {
		utils.Fail(c, 400, "参数错误")
		return
	}
	utils.Success(c, gin.H{
		"songs": songs,
	})
}

func GetArtistToplist(c *gin.Context) {
	typeStr := c.DefaultQuery("type", "0")
	typeInt, _ := strconv.Atoi(typeStr)
	if typeInt < 0 || typeInt > 2 {
		typeInt = 0
	}

	toplist := services.GetArtistToplist(typeInt)
	if toplist == nil {
		utils.Fail(c, 400, "参数错误")
		return
	}
	utils.Success(c, gin.H{
		"list": gin.H{
			"artists": toplist,
			"type":    typeInt,
		},
	})
}
