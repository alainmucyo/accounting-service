package company

import (
	"accounting-service/core/services/company"
	"accounting-service/events/dtos"
	companyEntity "accounting-service/store/entities/company"
	"accounting-service/store/kafka/producer"
	"encoding/json"
)

type EventHandler struct {
	service  *company.Service
	producer *producer.Producer
}

func New(service *company.Service,
	producer *producer.Producer) *EventHandler {
	return &EventHandler{service: service, producer: producer}
}

type response struct {
	Message string      `json:"message,omitempty"`
	Status  string      `json:"status,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

func (c *EventHandler) HandleCompanyCreateEvent(message []byte, _ string) {
	var value dtos.CompanyCreateDTO
	err := json.Unmarshal(message, &value)
	if err != nil {
		return
	}

	companyCreateObj := companyEntity.Company{
		CompanyID:        value.ID,
		AvailableBalance: value.Balance,
		ActualBalance:    value.Balance,
	}
	// Verify if company is already registered
	_, err = c.service.FindByCompanyID(value.ID)
	if err == nil {
		// TODO: Will find a better way to handle errors
		c.producer.Produce(response{
			Error:   "Company with ID " + companyCreateObj.CompanyID + " already registered",
			Message: "Company registration failed!",
			Status:  "FAILED",
		}, "accounting.company-registered.response")
		return
	}

	create, err := c.service.Create(companyCreateObj)
	if err != nil {
		print("Unable to register company")
		c.producer.Produce(response{
			Error:   err,
			Message: "Company registration failed!",
			Status:  "FAILED",
		}, "accounting.company-registered.response")
		return
	}

	c.producer.Produce(response{Data: create,
		Message: "Company registered successfully!",
		Status:  "SUCCESS",
	}, "accounting.company-registered.response")
}
