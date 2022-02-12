package channel

import (
	"accounting-service/store/entities"
)

type Channel struct {
	entities.Base
	Name string `gorm:"type:varchar(50)" json:"name"`
}

func (Channel) TableName() string {
	return "channels"
}
