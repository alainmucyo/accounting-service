package company

import (
	"accounting-service/store/entities/company"
	"accounting-service/store/postgres"
	"errors"
)

type Service struct {
	db *postgres.Database
}

func New(db *postgres.Database) *Service {
	return &Service{db: db}
}

func (s *Service) FindAll() ([]company.Company, error) {
	transactions := make([]company.Company, 0)
	if s.db.DB.Find(&transactions).Error != nil {
		return nil, errors.New("error while getting all companies")
	}
	return transactions, nil
}

func (s *Service) FindByCompanyID(companyID string) (company.Company, error) {
	var foundCompany company.Company
	if s.db.DB.Where(&company.Company{CompanyID: companyID}).First(&foundCompany).Error != nil {
		return company.Company{}, errors.New("company not found")
	}
	return foundCompany, nil
}

// FindByPK Finds a company by primary key
func (s *Service) FindByPK(companyID string) (company.Company, error) {
	var foundCompany company.Company
	if s.db.DB.Where(&company.Company{ID: companyID}).First(&foundCompany).Error != nil {
		return company.Company{}, errors.New("company not found")
	}
	return foundCompany, nil
}

func (s *Service) UpdateActualBalance(companyID string, balance int64) error {
	foundCompany, err := s.FindByPK(companyID)
	if err != nil {
		return err
	}
	foundCompany.ActualBalance = balance
	if s.db.DB.Model(&foundCompany).Where("id = ?", companyID).Updates(foundCompany).Error != nil {
		return errors.New("company not found")
	}
	return nil
}

func (s *Service) UpdateAvailableBalance(companyID string, balance int64) error {
	foundCompany, err := s.FindByPK(companyID)
	if err != nil {
		return err
	}
	foundCompany.AvailableBalance = balance
	if s.db.DB.Model(&foundCompany).Where("id = ?", companyID).Updates(foundCompany).Error != nil {
		return errors.New("company not found")
	}
	return nil
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
