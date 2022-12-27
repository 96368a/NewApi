package controller

import (
	"github.com/96368a/NewApi/services"
	"github.com/96368a/NewApi/utils"
	"github.com/gin-gonic/gin"
	"strconv"
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
