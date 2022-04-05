package user

import (
	"ginWeb/model"
	"ginWeb/service"
	"github.com/gin-gonic/gin"
	"log"
)

func PostInsertUserHandler(c *gin.Context) {
	var user model.User

	// 参数绑定
	if err := c.ShouldBind(&user); err == nil {
		log.Println(user)
		service.InsertUser(user)
	} else {
		log.Fatal("failure")
	}
}
