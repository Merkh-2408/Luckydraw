package router

import (
	"TESTGIN/controllers"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	//先生成一个实例
	r := gin.Default()
	r.GET("/", controllers.IndexController{}.ShowIndex)
	r.POST("/index", controllers.IndexController{}.GetIndex)

	return r
}
