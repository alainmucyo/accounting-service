package push

import (
	"accounting-service/api/dtos"
	"accounting-service/core/environment"
	"accounting-service/core/services/channel"
	"accounting-service/core/services/company"
	"accounting-service/core/services/transaction"
	"accounting-service/core/uuid"
	transactionEntity "accounting-service/store/entities/transaction"
	"accounting-service/store/kafka/producer"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	env                *environment.Environment
	transactionService *transaction.Service
	channelService     *channel.Service
	companyService     *company.Service
	uuidGenerator      *uuid.Generator
	kafkaProducer      *producer.Producer
}

func New(
	env *environment.Environment,
	transactionService *transaction.Service,
	channelService *channel.Service,
	companyService *company.Service,
	uuidGenerator *uuid.Generator,
	kafkaProducer *producer.Producer,
) *Handler {
	return &Handler{
		env:                env,
		transactionService: transactionService,
		channelService:     channelService,
		companyService:     companyService,
		uuidGenerator:      uuidGenerator,
		kafkaProducer:      kafkaProducer,
	}
}

type response struct {
	TransactionReference string `json:"transactionReference,omitempty"`
	GatewayReference     string `json:"gatewayReference,omitempty"`
}

func (h *Handler) HandleTransactionPushRequest(c *gin.Context) {
	var reqJSON dtos.TransactionDTO
	if err := c.BindJSON(&reqJSON); err != nil {
		println(err.Error())
		c.JSON(400, gin.H{"message": "Invalid request object"})
		return
	}
	// Validates transactions request
	err := reqJSON.Validate()
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	// Check if referenceID is already used
	_, err = h.transactionService.FindByReferenceID(reqJSON.TransactionReference)

	// If error is null, transaction with that reference is found and it is already used
	if err == nil {
		c.JSON(400, gin.H{"message": "Invalid request object"})
		return
	}

	// Find the company the request belongs to
	foundCompany, err := h.companyService.FindByCompanyID(reqJSON.CompanyId)

	// Error if company not found
	if err != nil {
		c.JSON(400, gin.H{"message": "Company with ID " + reqJSON.CompanyId + " not found"})
		return
	}

	// Check if the company has enough balance
	if foundCompany.AvailableBalance < reqJSON.Amount {
		c.JSON(400, gin.H{"message": "Not enough balance"})
		return
	}

	foundChannel, err := h.channelService.FindByName(reqJSON.Channel)
	if err != nil {
		c.JSON(400, gin.H{"message": "Channel " + reqJSON.Channel + " not found"})
		return
	}
	// If it reaches here, balance is okay. It will update available balance
	h.companyService.UpdateAvailableBalance(
		reqJSON.CompanyId,
		foundCompany.AvailableBalance-reqJSON.Amount,
	)
	transactionRequest := transactionEntity.Transaction{
		TransactionReference: reqJSON.TransactionReference,
		Amount:               reqJSON.Amount,
		Channel:              foundChannel,
		CompanyID:            reqJSON.CompanyId,
		GatewayReference:     h.uuidGenerator.GenerateUUID(), // Generates a UUID
		Msisdn:               reqJSON.Msisdn,
		Type:                 "push",
		GatewayStatus:        "pending",
	}
	createdTransaction, err := h.transactionService.Create(transactionRequest)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Something went wrong while creating a transaction",
		})
		return
	}
	switch foundChannel.Name {
	case "mtn-momo":
		// TODO: Sends a request to MTN MOMO
		break
	case "airtel-money":
		// TODO: sends to Airtel money
		break
	case "bk":
		// TODO: Sends to Airtel money
	}

	// TODO: This will be moved to a worker and in case a transaction failed, it will be handled there
	/*	// Transaction succeed
		transactionRequest = transactionEntity.Transaction{
			TransactionReference: reqJSON.TransactionReference,
			Amount:               reqJSON.Amount,
			Channel:              foundChannel,
			CompanyID:            reqJSON.CompanyId,
			GatewayReference:     h.uuidGenerator.GenerateUUID(), // Generates a UUID
			Msisdn:               reqJSON.Msisdn,
			Type:                 "push",
			GatewayStatus:        "success",
		}

		// TODO: This also will be moved to a worker. Here it is deducting but it will also refund.
		// Updates actual balance because transaction succeed
		h.companyService.UpdateActualBalance(
			reqJSON.CompanyId,
			foundCompany.AvailableBalance-reqJSON.Amount,
		)
	   // TODO: Update the transaction. */

	c.JSON(200, gin.H{
		"message": "Request received successfully",
		"data": response{
			TransactionReference: createdTransaction.TransactionReference,
			GatewayReference:     createdTransaction.GatewayReference,
		},
	})
	return
}
