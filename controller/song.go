package controller

import (
	"fmt"
	"github.com/96368a/NewApi/services"
	"github.com/96368a/NewApi/utils"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

func Search(c *gin.Context) {
	keywords := c.Query("keywords")
	limitStr := c.DefaultQuery("limit", "30")
	offsetStr := c.DefaultQuery("offset", "0")
	type1Str := c.DefaultQuery("type", "1")
	if keywords == "" {
		utils.Fail(c, 400, "参数错误")
		return
	}
	limit, _ := strconv.Atoi(limitStr)
	offset, _ := strconv.Atoi(offsetStr)
	type1, _ := strconv.Atoi(type1Str)

	switch type1 {
	case 10:
		{
			albums := services.SearchAlbum(keywords, int64(limit), int64(offset))
			utils.Success(c, gin.H{
				"result": gin.H{
					"albums": albums,
				},
			})
			return
		}
	case 100:
		{
			aritsts := services.SearchArtist(keywords, int64(limit), int64(offset))
			utils.Success(c, gin.H{
				"result": gin.H{
					"artists": aritsts,
				},
			})
		}
	case 1000:
		{
			playlists := services.SearchPlaylist(keywords, int64(limit), int64(offset))
			utils.Success(c, gin.H{
				"result": gin.H{
					"playlists": playlists,
				},
			})
		}

	default:
		{
			songs := services.SearchSong(keywords, int64(limit), int64(offset))
			utils.Success(c, gin.H{
				"result": gin.H{
					"songs": songs,
				},
			})
			return
		}
	}
}

func SongDetail(c *gin.Context) {
	idStr := c.Query("ids")
	idd := strings.Split(idStr, ",")
	var ids []int64
	for _, v := range idd {
		id, err := strconv.Atoi(v)
		if err == nil {
			ids = append(ids, int64(id))
		}
	}
	if ids == nil {
		utils.Fail(c, 400, "参数错误")
		return
	}
	songs := services.GetSongs(ids)

	if songs == nil {
		utils.Success(c, gin.H{
			"songs": nil,
		})
		return
	}
	fmt.Printf("%v", idd)
	utils.Success(c, gin.H{
		"songs": songs,
	})
}
