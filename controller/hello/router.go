package hello

import "github.com/gin-gonic/gin"

func Router(e *gin.Engine) {
	e.GET("/hello/say", GetHelloHandler)
}
