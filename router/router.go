package router

import (
	"TESTGIN/controllers"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	//先生成一个实例
	r := gin.Default()
	r.LoadHTMLGlob("templates/*") // 加载模板文件
	indexController := controllers.IndexController{}
	lotteryController := controllers.LotteryController{}
	r.GET("/", controllers.IndexController{}.ShowIndex)
	r.POST("/index", indexController.GetIndex)
	r.POST("/lottery", lotteryController.Draw) // 添加对 /lottery 路径的处理
	return r
}
