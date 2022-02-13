package channel

import (
	"accounting-service/store/entities"
)

type Channel struct {
	ID string `sql:"type:uuid;primary_key;default:uuid_generate_v4()"`
	entities.Base
	Name string `gorm:"type:varchar(50)" json:"name"`
}

func (Channel) TableName() string {
	return "channels"
}
