package dtos

type TransactionDTO struct {
	CompanyId            string `json:"company_id,omitempty"`
	TransactionType      string `json:"transaction_type,omitempty"`
	Amount               int64  `json:"amount,omitempty"`
	Msisdn               string `json:"msisdn,omitempty"`
	Channel              string `json:"channel,omitempty"`
	TransactionReference string `json:"transaction_reference,omitempty"`
}
