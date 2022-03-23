package dtos

type CompanyCreateDTO struct {
	ID      string `json:"id,omitempty"`
	Balance int64  `json:"balance,omitempty"`
}
