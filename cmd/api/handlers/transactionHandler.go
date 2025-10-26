package handlers

import (
	"bucketWise/pkg/domain"
	"bucketWise/pkg/dto"
	"bucketWise/pkg/ports"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	TransactionUC ports.TransactionUseCase
}

// CreateTransaction godoc
// @Summary Create a new Transaction
// @Description Creates a new Transaction in the system
// @Tags transactions
// @Accept json
// @Produce json
// @Param Transaction body dto.TransactionCreateRequest true "Transaction data with amount, description and date"
// @Success 200 {object} dto.TransactionResponse
// @Failure 400 {object} map[string]string
// @Failure 409 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /transactions [post]
func (h TransactionHandler) CreateTransaction(c *gin.Context) {
	var req dto.TransactionCreateRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tx := domain.Transaction{
		Amount:      req.Amount,
		Description: req.Description,
		Date:        req.Date,
	}

	insertedTx, err := h.TransactionUC.CreateTransactionUseCase(tx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	resp := dto.TransactionResponse{
		ID:           insertedTx.ID,
		Date:         insertedTx.Date,
		Amount:       insertedTx.Amount,
		Description:  insertedTx.Description,
		CategoryID:   insertedTx.CategoryID,
		CategoryName: insertedTx.CategoryName,
		Type:         string(insertedTx.Type),
	}
	c.JSON(http.StatusCreated, resp)
}
