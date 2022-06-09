package user

import "github.com/gin-gonic/gin"

func Router(e *gin.Engine) {
	e.POST("/user/insert", InsertUserHandler)
	e.POST("/user/login", LoginHandler)
}
