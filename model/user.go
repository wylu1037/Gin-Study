package model

import (
	"ginWeb/config/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

// User Entity-用户
type User struct {
	gorm.Model
	Account  string
	Password string
	Type     int
	UserName string
}

// InsertUser ModelSql-插入用户
func InsertUser(user User) uint {
	result := mysql.DB.Create(&user)
	err := result.Error
	affected := result.RowsAffected
	if err != nil {
		log.Fatalf("Insert user failure, err: %v", err)
		return 0
	} else {
		id := user.ID
		log.Printf("Insert user success, %d rows affedted, return primary key id = %d", affected, id)
		return id
	}
}

// ExistDuplicatedName 校验是否存在重名
func ExistDuplicatedName(name string) bool {
	var count int8
	mysql.DB.Model(&User{}).Where("user_name = ?", name).Limit(1).Count(&count)
	return count > 0
}

// FindUser ModelSql-查找用户
func FindUser(userName, password string) *User {
	user := User{}
	mysql.DB.Where("user_name = ? and password = ?", userName, password).Find(&user)
	return &user
}
