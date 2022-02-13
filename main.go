package main

import (
	transactionHandler "accounting-service/api/handlers/transaction"
	"accounting-service/core/environment"
	"accounting-service/core/services/channel"
	"accounting-service/core/services/company"
	"accounting-service/core/services/transaction"
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

	// FYI: Will look for better way to handle dependency injection in Go
	cache := redis.New(envs, ctx)
	db := postgres.New(envs)

	db.Migrate()

	channelService := channel.New(db)
	transactionService := transaction.New(cache, db)
	companyService := company.New(db)
	transactionController := transactionHandler.New(envs, transactionService, channelService, companyService)
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/transaction/request", transactionController.HandleTransactionRequest)
	r.Static("/api-docs", "./swaggerui")
	r.Run(":3000")
}
