package hello

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetHelloHandler(c *gin.Context) {
	// 获取query参数
	msg := c.DefaultQuery("msg", "Hello Go!")
	c.JSON(http.StatusOK, gin.H{
		"method": "SayHello",
		"data":   msg,
	})
}
