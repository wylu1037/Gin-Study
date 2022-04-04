package main

import (
	"fmt"
	"ginWeb/conf"
	"ginWeb/conf/mysql"
	"ginWeb/conf/redis"
	"ginWeb/controller/hello"
	"ginWeb/router"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// 先于main函数执行，用于配置文件读取，数据库连接初始化等操作
func init() {
	// 读取配置文件
	conf.ReadProps()
	// 初始化mysql连接
	mysql.InitConnect()
	// 初始化redis连接
	err := redis.InitConnect()
	if err != nil {
		log.Fatalf("redis initialize failed, err: %v", err)
	}
}

// 启动类入口
// @title Golang Gin API
// @version 1.0
// @description An example of gin
// @termsOfService https://gitee.com/luwy5180/ginWeb
// @license.name license
// @license.url https://gitee.com/luwy5180/ginWeb
func main() {
	gin.SetMode(conf.ServerSetting.RunMode)

	// 加载路由
	router.Include(hello.Router)
	routerInit := router.Init()

	readTimeout := conf.ServerSetting.ReadTimeout
	writeTimeout := conf.ServerSetting.WriteTimeout
	port := conf.ServerSetting.HttpPort
	endPoint := fmt.Sprintf(":%d", port)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routerInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("main start failed, err: %v", err)
	}

	/*err := r.Run()
	if err != nil {
		log.Fatalf("gin web service startup failed, err: %v \n\n", err)
	}*/
}
