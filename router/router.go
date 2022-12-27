package router

import (
	"github.com/96368a/NewApi/controller"
	"github.com/96368a/NewApi/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	//跨域处理
	r.Use(middleware.CORSMiddleware())

	r.GET("/login", controller.LoginByEmail)
	r.POST("/login", controller.LoginByEmail)
	r.GET("/login/cellphone", controller.LoginByPhone)
	r.POST("/login/cellphone", controller.LoginByPhone)
	r.GET("/login/status", controller.LoginStatus)
	r.GET("/search", controller.Search)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"code": 404,
			"msg":  "页面不存在",
		})
	})

	return r
}
