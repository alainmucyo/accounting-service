package company

import (
	"accounting-service/store/entities/company"
	"accounting-service/store/postgres"
	"errors"
)

type Service struct {
	db    *postgres.Database
}

func New(db *postgres.Database) *Service {
	return &Service{db: db}
}

func (s *Service) FindAll() ([]company.Company, error) {
	apps := make([]company.Company, 0)
	if s.db.DB.Find(&apps).Error != nil {
		return nil, errors.New("error while getting all apps")
	}
	return apps, nil
}

func (s *Service) Create(company company.Company) (company.Company, error) {
	ctx := s.db.DB.Begin()
	if ctx.Error != nil {
		return company, errors.New("error start")
	}
	if s.db.DB.Save(&company).Error != nil {
		ctx.Rollback()
		return company, errors.New("error save")
	}
	if ctx.Commit().Error != nil {
		ctx.Rollback()
		return company, errors.New("error commit")
	}
	return company, nil
}
