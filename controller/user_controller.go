package controller

import (
	"ginWeb/model"
	"ginWeb/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func UserRouter(e *gin.Engine) {
	e.POST("/user/insert", insertUserHandler)
	e.POST("/user/login", loginHandler)
}

func insertUserHandler(c *gin.Context) {
	var user model.User

	// 参数绑定
	if err := c.ShouldBind(&user); err == nil {
		log.Println(user)
		service.InsertUser(user)
	} else {
		log.Fatal("failure")
	}
}

func loginHandler(c *gin.Context) {
	userName := c.PostForm("userName")
	password := c.PostForm("password")
	token := service.Login(userName, password)

	if token != "" {
		c.JSON(http.StatusOK, gin.H{
			"token":   token,
			"message": "登录成功",
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"token":   token,
			"message": "登录失败",
		})
	}
}
