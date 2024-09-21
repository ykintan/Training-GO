package main

import (
	"log"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	time.Sleep(3 * time.Second)
	log.Println("Hello Docker Compose!")
	dsn := "postgresql://postgres:P4ssw0rd@pg-db:5432/go_hello"
	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Hello world from PostgreSQL!")

	sqlDB, err := gormDB.DB()
	if err != nil {
		log.Fatalln(err)
	}

	if err := sqlDB.Ping(); err != nil {
		log.Fatalln(err)
	}

	log.Println("Connected to database")

	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis-db:6379",
		Password: "redispass",
		DB:       0,
	})
	log.Println(rdb)
}
