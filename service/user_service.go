package service

import (
	"ginWeb/model"
	"log"
)

func InsertUser(user model.User) {
	// 处理业务逻辑：字段校验、重名校验
	exist := model.ExistDuplicatedName(user.UserName)
	if exist {
		log.Printf("名称%s重复，请更换后重试 \n", user.UserName)
		return
	}
	model.InsertUser(user)
}
