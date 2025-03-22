package controllers

import (
	"TESTGIN/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IndexController struct{}

func (u IndexController) GetIndex(c *gin.Context) {
	utils.ReturnSuccess(c, 200, "成功", "这个是一个成功的信息", 1)
}

func (u IndexController) ShowIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}


  