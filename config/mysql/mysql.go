package mysql

import (
	"fmt"
	"ginWeb/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var DB *gorm.DB

// InitConnect 初始化MySql连接
func InitConnect() {
	var err error
	dialect := config.DatabaseSetting.Dialect
	url := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DatabaseSetting.User,
		config.DatabaseSetting.Password,
		config.DatabaseSetting.Host,
		config.DatabaseSetting.DatabaseName)
	DB, err = gorm.Open(dialect, url)
	if err != nil {
		log.Fatalf("gorm open mysql failed, err: %v", err)
		return
	}

	DB.SingularTable(true)
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)
}

func CloseMysqlDB() {
	err := DB.Close()
	if err != nil {
		log.Fatalf("mysql close connect failed, err: %v \n", err)
	}
}
