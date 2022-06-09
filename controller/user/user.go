package user

import (
	"ginWeb/model"
	"ginWeb/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// InsertUserHandler Controller-POST-插入用户记录
func InsertUserHandler(c *gin.Context) {
	var user model.User

	// 参数绑定
	if err := c.ShouldBind(&user); err == nil {
		log.Println(user)
		service.InsertUser(user)
	} else {
		log.Fatal("failure")
	}
}

// LoginHandler Controller-POST-用户登录
func LoginHandler(c *gin.Context) {
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
