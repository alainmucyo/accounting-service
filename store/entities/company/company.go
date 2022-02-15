package company

import "accounting-service/store/entities"

type Company struct {
	ID               string `sql:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	CompanyID        string `gorm:"type:varchar(50)" json:"company_id"`
	ActualBalance    int64  `gorm:"default:0" json:"actual_balance"`
	AvailableBalance int64  `gorm:"default:0" json:"available_balance"`
	entities.Base
}

func (Company) TableName() string {
	return "companies"
}
