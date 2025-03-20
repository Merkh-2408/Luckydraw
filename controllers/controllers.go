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

type LotteryController struct{}

func (lc LotteryController) Draw(c *gin.Context) {
	items := c.PostForm("items")
	result, err := utils.DrawFromText(items)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": result})
}
