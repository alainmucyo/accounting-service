package company

import (
	"accounting-service/store/entities"
)

type Company struct {
	entities.Base
	CompanyID string `gorm:"type:varchar(50)" json:"company_id"`
	Balance   int64  `gorm:"default:0" json:"balance"`
}

func (Company) TableName() string {
	return "companies"
}
