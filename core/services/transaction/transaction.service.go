package transaction

import (
	"accounting-service/store/entities/transaction"
	"accounting-service/store/postgres"
	"accounting-service/store/redis"
	"errors"
)

type Service struct {
	cache *redis.Cache
	db    *postgres.Database
}

func New(cache *redis.Cache, db *postgres.Database) *Service {
	return &Service{cache: cache, db: db}
}

func (s *Service) FindAll() ([]transaction.Transaction, error) {
	apps := make([]transaction.Transaction, 0)
	if s.db.DB.Find(&apps).Error != nil {
		return nil, errors.New("error while getting all apps")
	}
	return apps, nil
}

func (s *Service) Create(transaction transaction.Transaction) (transaction.Transaction, error) {
	ctx := s.db.DB.Begin()
	if ctx.Error != nil {
		return transaction, errors.New("error start")
	}
	if s.db.DB.Save(&transaction).Error != nil {
		ctx.Rollback()
		return transaction, errors.New("error save")
	}
	if ctx.Commit().Error != nil {
		ctx.Rollback()
		return transaction, errors.New("error commit")
	}
	return transaction, nil
}
