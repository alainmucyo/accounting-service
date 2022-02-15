package channel

import (
	"accounting-service/store/entities"
)

type Channel struct {
	ID   string `sql:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Name string `gorm:"type:varchar(50)" json:"name"`
	entities.Base
}

func (Channel) TableName() string {
	return "channels"
}
