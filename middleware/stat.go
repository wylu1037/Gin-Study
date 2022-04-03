package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

// RequestConsume 定义统计请求耗时的中间件
func RequestConsume() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		cost := time.Since(start)
		uri := c.Request.RequestURI
		method := c.Request.Method
		log.Printf("receive request: uri = %s, method = %s, time cost = %v",
			uri, method, cost)
	}
}
