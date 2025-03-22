package controllers

import (
	"TESTGIN/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

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