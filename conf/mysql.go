package conf

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var DB *gorm.DB

// InitMysqlConnect 初始化MySql连接
func InitMysqlConnect() {
	var err error
	dialect := DatabaseSetting.Dialect
	url := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		DatabaseSetting.User,
		DatabaseSetting.Password,
		DatabaseSetting.Host,
		DatabaseSetting.DatabaseName)
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
