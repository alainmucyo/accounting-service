package main

import (
	"accounting-service/api/handlers/transaction/pull"
	"accounting-service/api/handlers/transaction/push"
	"accounting-service/core/environment"
	"accounting-service/core/services/channel"
	"accounting-service/core/services/company"
	"accounting-service/core/services/transaction"
	"accounting-service/core/uuid"
	companyEvents "accounting-service/events/handlers/company"
	"accounting-service/store/kafka/consumer"
	"accounting-service/store/kafka/producer"
	"accounting-service/store/kafka/topics"
	"accounting-service/store/postgres"
	"accounting-service/store/redis"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

// TODO: This main.go file will be moved to cmd/accounting
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// Getting app context
	ctx := context.Background()

	// Initializes available env
	envs := environment.New(
		os.Getenv("PORT"),
		os.Getenv("REDIS_URL"),
		os.Getenv("REDIS_PASSWORD"),
		os.Getenv("DB_URL"),
		os.Getenv("KAFKA_BROKER_URL"),
		os.Getenv("KAFKA_GROUP_ID"),
	)

	// FYI: Will look for better way to handle dependency injection in Go
	cache := redis.New(envs, ctx)
	db := postgres.New(envs)
	kafkaProducer := producer.New(envs)

	// Migrates databases if they are not available
	db.Migrate()

	// Initialize services
	channelService := channel.New(db)

	// This creates one channel names mtn-momo
	channelService.Seed()

	transactionService := transaction.New(cache, db)
	companyService := company.New(db)

	// Kafka company event handler. This handles all Kafka requests related to companies
	companyEventHandler := companyEvents.New(companyService, kafkaProducer)
	kafkaTopics := topics.New(envs, companyEventHandler)
	kafkaConsumer := consumer.New(envs, kafkaTopics)

	/**
	Consumer blocks, it's like an endless loop that is waiting for messages.
	It is being initialized in Go routine for not to block everything else
	*/
	go kafkaConsumer.Consume()

	// Initialise a UUID generator.
	uuidGenerator := uuid.New()

	// Transaction push controller
	pushController := push.New(
		envs,
		transactionService,
		channelService,
		companyService,
		uuidGenerator,
		kafkaProducer,
	)

	// Transaction pull controller
	pullController := pull.New(
		envs,
		transactionService,
		channelService,
		companyService,
		uuidGenerator,
		kafkaProducer,
	)

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/api/v1/payment/push", pushController.HandleTransactionPushRequest)
	r.POST("/api/v1/payment/pull", pullController.HandleTransactionPullRequest)

	r.Static("/api-docs", "./swaggerui")
	r.Run(":" + envs.Port)
}
