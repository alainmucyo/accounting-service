package transaction

import (
	"accounting-service/core/environment"
	"accounting-service/core/services/channel"
	"accounting-service/core/services/company"
	"accounting-service/core/services/transaction"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	env                *environment.Environment
	transactionService *transaction.Service
	channelService     *channel.Service
	companyService     *company.Service
}

func New(
	env *environment.Environment,
	transactionService *transaction.Service,
	channelService *channel.Service,
	companyService *company.Service,
) *Handler {
	return &Handler{
		env:                env,
		transactionService: transactionService,
		channelService:     channelService,
		companyService:     companyService,
	}
}

func (h *Handler) HandleTransactionRequest(c *gin.Context) {
	//TODO: Handle transaction request
}
