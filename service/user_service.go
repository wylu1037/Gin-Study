package service

import (
	"ginWeb/config/jwt"
	"ginWeb/model"
	"log"
)

// InsertUser Service-插入用户
func InsertUser(user model.User) {
	// 处理业务逻辑：字段校验、重名校验
	exist := model.ExistDuplicatedName(user.UserName)
	if exist {
		log.Printf("名称%s重复，请更换后重试 \n", user.UserName)
		return
	}
	model.InsertUser(user)
}

// Login Service-用户登录
func Login(userName, password string) string {
	user := model.FindUser(userName, password)
	if user.ID != 0 {
		token, _ := jwt.CreateToken(userName, user.ID)
		return token
	} else {
		return ""
	}
}
