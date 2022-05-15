package config

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

const (
	FilePath        = "config/application.ini"
	SectionDatabase = "database"
	SectionRedis    = "redis"
	SectionServer   = "server"
)

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}

type Database struct {
	Dialect      string
	User         string
	Password     string
	Host         string
	DatabaseName string
	TablePrefix  string
}

var DatabaseSetting = &Database{}

type Redis struct {
	Host        string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

var RedisSetting = &Redis{}

var cfg *ini.File

// ReadProps 读取配置
func ReadProps() {
	var err error
	cfg, err = ini.Load(FilePath)
	if err != nil {
		log.Fatalf("setting read file application.ini failed, err: %v \n", err)
	}
	mapTo(SectionDatabase, DatabaseSetting)
	mapTo(SectionRedis, RedisSetting)
	mapTo(SectionServer, ServerSetting)
}

// 将配置文件的属性匹配赋值到结构体上
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("setting mapTo %s failed, err: %v", section, err)
	}
}
