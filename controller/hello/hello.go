package hello

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetHelloHandler
// @Summary SayHello接口
// @Tags Hello相关接口
// @Router /hello/say [get]
func GetHelloHandler(c *gin.Context) {
	// 获取query参数
	msg := c.DefaultQuery("msg", "Hello Go!")
	c.JSON(http.StatusOK, gin.H{
		"method": "SayHello",
		"data":   msg,
	})
}
