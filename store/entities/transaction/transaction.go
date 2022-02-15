package transaction

import (
	"accounting-service/store/entities"
	"accounting-service/store/entities/channel"
)

type Transaction struct {
	ID                   string          `sql:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	CompanyID            string          `gorm:"type:varchar(50)" json:"company_id"`
	Amount               int64           `json:"amount"`
	Msisdn               string          `gorm:"type:varchar(20)" json:"msisdn"`
	ChannelId            string          `json:"channel_id"`
	Channel              channel.Channel `gorm:"foreignKey:ChannelId;references:ID" json:"channel"`
	Type                 string          `gorm:"type:varchar(50),default:push" json:"transaction_type"`
	TransactionReference string          `gorm:"type:varchar(50)" json:"transaction_reference"`
	GatewayReference     string          `gorm:"type:varchar(50)" json:"gateway_reference"`
	GatewayStatus        string          `gorm:"type:varchar(50)" json:"gateway_status"`
	GatewayStatusCode    int             `json:"gateway_status_code"`
	entities.Base
}

func (Transaction) TableName() string {
	return "transactions"
}
