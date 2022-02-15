package dtos

import (
	"errors"
	"regexp"
)

type TransactionDTO struct {
	CompanyId            string `json:"company_id,omitempty"`
	Amount               int64  `json:"amount,omitempty"`
	Msisdn               string `json:"msisdn,omitempty"`
	Channel              string `json:"channel,omitempty"`
	TransactionReference string `json:"transaction_reference,omitempty"`
}

func (t TransactionDTO) Validate() error {
	r, _ := regexp.Compile("(^(07[8,9])[0-9]{7}$)")
	if !r.MatchString(t.Msisdn) {
		return errors.New("invalid phone msisdn")
	}
	if t.Amount < 10 {
		return errors.New("amount should not be less than 10")
	}
	if t.CompanyId == "" {
		return errors.New("company ID is required")
	}
	if t.Channel == "" {
		return errors.New("channel is required")
	}
	if t.TransactionReference == "" {
		return errors.New("transaction reference is required")
	}
	return nil
}
