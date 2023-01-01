package main

import (
	"github.com/96368a/NewApi/router"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	gin.SetMode(gin.DebugMode)
	r := router.InitRouter()
	r.SetTrustedProxies(nil)

	//创建资源文件夹
	_, err := os.Lstat("resources/musics")
	if os.IsNotExist(err) {
		os.Mkdir("resources/musics", 0755)
	}
	_, err = os.Lstat("resources/images")
	if os.IsNotExist(err) {
		os.Mkdir("resources/images", 0755)
	}

	panic(r.Run(":8800"))
}
