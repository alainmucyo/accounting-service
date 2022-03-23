package postgres

import (
	"accounting-service/core/environment"
	"accounting-service/store/entities/channel"
	"accounting-service/store/entities/company"
	"accounting-service/store/entities/transaction"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

type Database struct {
	env *environment.Environment
	DB  *gorm.DB
}

func New(env *environment.Environment) *Database {
	var err error
	// Use postgres DB URL from .env
	db, err := gorm.Open(postgres.Open(env.DBURL), &gorm.Config{})
	if err != nil {
		os.Exit(2)
	}
	return &Database{DB: db, env: env}
}

// Migrate Auto migrates tables
func (db Database) Migrate() {
	err := db.DB.AutoMigrate(company.Company{}, channel.Channel{}, transaction.Transaction{})
	if err != nil {
		fmt.Println("Failed to migrate tables")
		fmt.Println(err.Error())
		os.Exit(2)
	}
}
