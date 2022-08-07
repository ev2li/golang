package models

import (
	"fmt"

	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB = Init()
var RDB = InitRedisDB()

func Init() *gorm.DB {
	dsn := "root:123456@tcp(127.0.0.1:3306)/online?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("gorm init error:", err)
	}
	return db
}

func InitRedisDB() *redis.Client {
	var rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return rdb
}
