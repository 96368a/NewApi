package controller

import (
	"github.com/96368a/NewApi/services"
	"github.com/96368a/NewApi/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetPlaylistToplist(c *gin.Context) {
	toplist := services.GetPlaylistToplist()
	utils.Success(c, gin.H{
		"list": toplist,
	})

}

func GetPlaylistPersonalized(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "30")
	limit, _ := strconv.Atoi(limitStr)
	if limit <= 0 {
		limit = 30
	}
	playlists := services.GetPlaylistPersonalized(int64(limit))
	var playlistDto []map[string]any
	for _, playlist := range playlists {
		playlistDto = append(playlistDto, map[string]any{
			"id":                    playlist.ID,
			"name":                  playlist.Name,
			"picUrl":                playlist.CoverImgURL,
			"playCount":             playlist.PlayCount,
			"trackCount":            playlist.TrackCount,
			"trackNumberUpdateTime": playlist.TrackUpdateTime,
			"type":                  playlist.AdType,
		})
	}
	utils.Success(c, gin.H{
		"result": playlistDto,
	})

}

func GetPlaylistDetail(c *gin.Context) {
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
	playlist := services.GetOnePlaylist(int64(id))
	utils.Success(c, gin.H{
		"playlist": playlist,
	})

}
