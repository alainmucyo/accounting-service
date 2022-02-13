package company

import (
	"accounting-service/store/entities"
)

type Company struct {
	entities.Base
	ID               string `sql:"type:uuid;primary_key;default:uuid_generate_v4()"`
	CompanyID        string `gorm:"type:varchar(50)" json:"company_id"`
	ActualBalance    int64  `gorm:"default:0" json:"actual_balance"`
	AvailableBalance int64  `gorm:"default:0" json:"available_balance"`
}

func (Company) TableName() string {
	return "companies"
}
