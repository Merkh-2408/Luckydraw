package utils

import (
	"math/rand"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type JsonSteuct struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data"`
	Count int64       `json:"count"`
}

func ReturnSuccess(c *gin.Context, code int, msg string, data interface{}, count int64) {
	json := &JsonSteuct{Code: code, Msg: msg, Data: data, Count: count}
	c.JSON(200, json)
}

func DrawFromText(text string) (string, error) {
	rand.Seed(time.Now().UnixNano())
	lines := strings.Split(text, "\n")
	items := make([]string, 0)
	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)
		if trimmedLine != "" {
			items = append(items, trimmedLine)
		}
	}
	if len(items) == 0 {
		return "", nil
	}
	return items[rand.Intn(len(items))], nil
}
