package transaction

import (
	"accounting-service/api/dtos"
	"accounting-service/core/environment"
	"accounting-service/core/services/channel"
	"accounting-service/core/services/company"
	"accounting-service/core/services/transaction"
	"accounting-service/core/uuid"
	transaction2 "accounting-service/store/entities/transaction"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	env                *environment.Environment
	transactionService *transaction.Service
	channelService     *channel.Service
	companyService     *company.Service
	uuidGenerator      *uuid.Generator
}

func New(
	env *environment.Environment,
	transactionService *transaction.Service,
	channelService *channel.Service,
	companyService *company.Service,
	uuidGenerator *uuid.Generator,
) *Handler {
	return &Handler{
		env:                env,
		transactionService: transactionService,
		channelService:     channelService,
		companyService:     companyService,
		uuidGenerator:      uuidGenerator,
	}
}

func (h *Handler) HandleTransactionRequest(c *gin.Context) {
	var reqJSON dtos.TransactionDTO
	if err := c.BindJSON(&reqJSON); err != nil {
		println(err.Error())
		c.JSON(400, gin.H{"message": "Invalid request object"})
		return
	}

	// Check if referenceID is already used
	_, err := h.transactionService.FindByReferenceID(reqJSON.TransactionReference)

	// If error is null, means reference to found, and it is already used
	if err == nil {
		c.JSON(400, gin.H{"message": "Invalid request object"})
		return
	}

	// Find the company the request belongs to
	company, err := h.companyService.FindByCompanyID(reqJSON.CompanyId)

	// Error if company not found
	if err != nil {
		// TODO: Create a company if not found
		c.JSON(400, gin.H{"message": "Company with ID " + reqJSON.CompanyId + " not found"})
		return
	}

	// Check if the company has enough balance
	if company.AvailableBalance < reqJSON.Amount {
		c.JSON(400, gin.H{"message": "Not enough balance"})
		return
	}

	channel, err := h.channelService.FindByName(reqJSON.Channel)
	if err != nil {
		// TODO: Create a company if not found
		c.JSON(400, gin.H{"message": "Channel " + reqJSON.Channel + " not found"})
		return
	}
	// If it reaches here, balance is okay. I will update available balance
	h.companyService.UpdateAvailableBalance(reqJSON.CompanyId, company.AvailableBalance-reqJSON.Amount)

	transaction := transaction2.Transaction{
		TransactionReference: reqJSON.TransactionReference,
		Amount:               reqJSON.Amount,
		Channel:              channel,
		CompanyID:            reqJSON.CompanyId,
		GatewayReference:     h.uuidGenerator.GenerateUUID(),
	}
	h.transactionService.Create(transaction)
	c.JSON(400, gin.H{"message": "Request received successfully"})
	return
}
