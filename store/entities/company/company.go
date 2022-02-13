package company

import (
	"accounting-service/store/entities"
)

type Company struct {
	entities.Base
	CompanyID        string `gorm:"type:varchar(50)" json:"company_id"`
	ActualBalance    int64  `gorm:"default:0" json:"actual_balance"`
	AvailableBalance int64  `gorm:"default:0" json:"available_balance"`
}

func (Company) TableName() string {
	return "companies"
}
