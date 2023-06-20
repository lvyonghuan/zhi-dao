package dao

import (
	"github.com/go-redis/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB
var Redis *redis.Client

func InitDB() {
	dsn := "root:42424242@tcp(127.0.0.1:3306)/zhidao"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("mysql初始化错误:%v", err)
	}
	DB = db
}

func InitRedis() {
	Redis = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		DB:       0,
		Password: "",
	})
	err := Redis.Ping().Err()
	if err != nil {
		log.Fatalf("redis初始化错误：%v", err)
	}
}
