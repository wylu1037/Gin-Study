package router

import (
	_ "ginWeb/docs"
	"ginWeb/middleware"
	"github.com/gin-gonic/gin"
	gs "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// Router 定义路由函数类型，泛指controller层的函数
type Router func(*gin.Engine)

// 定义路由数组，管理项目所有路由
var routes []Router

// Include 注册路由
func Include(route ...Router) {
	routes = append(routes, route...)
}

// Init 初始化路由
func Init() *gin.Engine {
	r := gin.Default()
	// 注册全局中间件
	r.Use(middleware.RequestConsume())

	for _, rou := range routes {
		rou(r)
	}
	//
	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

	return r
}
