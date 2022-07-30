package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HelloRouter(e *gin.Engine) {
	e.GET("/hello/say", sayHandler)
}

// sayHandler godoc
// @Summary SayHello接口
// @Description  get string
// @Tags Hello相关接口
// @Router /hello/say [get]
func sayHandler(c *gin.Context) {
	// 获取query参数
	msg := c.DefaultQuery("msg", "Hello Go!")
	c.JSON(http.StatusOK, gin.H{
		"method": "SayHello",
		"data":   msg,
	})
}
