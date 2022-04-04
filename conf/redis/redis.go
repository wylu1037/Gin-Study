package redis

import (
	"encoding/json"
	"ginWeb/conf"
	"github.com/gomodule/redigo/redis"
	"log"
	"time"
)

var RedisConnect *redis.Pool

// InitRedisConnect 初始化Redis连接
func InitRedisConnect() error {
	RedisConnect = &redis.Pool{
		MaxIdle:     conf.RedisSetting.MaxIdle,
		MaxActive:   conf.RedisSetting.MaxActive,
		IdleTimeout: conf.RedisSetting.IdleTimeout,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", conf.RedisSetting.Host)
			if err != nil {
				return nil, err
			}
			if conf.RedisSetting.Password != "" {
				_, err := conn.Do("AUTH", conf.RedisSetting.Password)
				if err != nil {
					conn.Close()
					return nil, err
				}
			}
			return conn, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
	return nil
}

// Select a database with the SELECT command:
func Select(db int) (redis.Conn, error) {
	conn := RedisConnect.Get()
	if _, err := conn.Do("SELECT", db); err != nil {
		return nil, err
	}
	return conn, nil
}

// Set a key/value
func Set(key string, data interface{}, time int, db int) error {
	conn, err := Select(db)
	defer conn.Close()
	if conn == nil || err != nil {
		log.Fatalf("get redis connection failed, err: %v", err)
		return err
	}

	value, err := json.Marshal(data)
	if err != nil {
		return err
	}
	_, err = conn.Do("SET", key, value)
	if err != nil {
		return err
	}

	if time > 0 {
		_, err = conn.Do("EXPIRE", key, time)
		if err != nil {
			return err
		}
	}

	return nil
}

// Exists check a key
func Exists(key string, db int) bool {
	conn, err := Select(db)
	defer conn.Close()
	if conn == nil || err != nil {
		log.Fatalf("get redis connection failed, err: %v", err)
		return false
	}

	exists, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return false
	}

	return exists
}

// Get a key
func Get(key string, db int) ([]byte, error) {
	conn, err := Select(db)
	defer conn.Close()
	if conn == nil || err != nil {
		log.Fatalf("get redis connection failed, err: %v", err)
		return nil, err
	}

	reply, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return nil, err
	}

	return reply, nil
}

// Delete delete a kye
func Delete(key string, db int) (bool, error) {
	conn, err := Select(db)
	defer conn.Close()
	if conn == nil || err != nil {
		log.Fatalf("get redis connection failed, err: %v", err)
		return false, err
	}

	return redis.Bool(conn.Do("DEL", key))
}

// LikeDeletes batch delete
func LikeDeletes(key string, db int) error {
	conn, err := Select(db)
	defer conn.Close()
	if conn == nil || err != nil {
		log.Fatalf("get redis connection failed, err: %v", err)
		return err
	}

	keys, err := redis.Strings(conn.Do("KEYS", "*"+key+"*"))
	if err != nil {
		return err
	}

	for _, key := range keys {
		_, err = Delete(key, db)
		if err != nil {
			return err
		}
	}

	return nil
}