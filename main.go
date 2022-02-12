package main

import (
	"accounting-service/core/environment"
	"accounting-service/store/postgres"
	"accounting-service/store/redis"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	ctx := context.Background()
	envs := environment.New(
		os.Getenv("PORT"),
		os.Getenv("REDIS_URL"),
		os.Getenv("REDIS_PASSWORD"),
		os.Getenv("DB_URL"),
	)

	_ = redis.New(envs, ctx)
	db := postgres.New(envs)

	db.Migrate()

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Static("/api-docs", "./swaggerui")
	r.Run(":3000")
}
