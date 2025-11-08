package handlers

import (
	"bucketWise/pkg/domain"
	"bucketWise/pkg/dto"
	"bucketWise/pkg/ports"
	"net/http"
	"strings"

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

// ListTransactions godoc
// @Summary Get transactions (optionally filtered by category)
// @Description Retrieves all transactions, or those matching a category if provided
// @Tags transactions
// @Param category query string false "Category name (optional)"
// @Produce json
// @Success 200 {array} dto.TransactionResponse "List of transactions"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /transactions [get]
func (h TransactionHandler) ListTransactions(c *gin.Context) {
	cat := strings.TrimSpace(c.Query("category"))
	transactionList, err := h.TransactionUC.ListTransactionsUseCase(cat)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var response []dto.TransactionResponse
	for _, tx := range transactionList {
		response = append(response, dto.TransactionResponse{
			ID:           tx.ID,
			Amount:       tx.Amount,
			Description:  tx.Description,
			Date:         tx.Date,
			CategoryID:   tx.CategoryID,
			CategoryName: tx.CategoryName,
			Type:         string(tx.Type),
		})
	}

	c.JSON(http.StatusOK, gin.H{"transactions": response})
}

// DeleteTransactions godoc
// @Summary Delete multiple transactions
// @Description Deletes one or more transactions by their IDs
// @Tags transactions
// @Accept  json
// @Produce  json
// @Param body body dto.TransactionsDeleteRequest true "Array of transaction IDs to delete"
// @Success 200 {object} map[string]interface{} "Deletion result"
// @Failure 400 {object} map[string]string "Invalid request"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /transactions [delete]
func (h TransactionHandler) DeleteTransactions(c *gin.Context) {
	var req dto.TransactionsDeleteRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	if len(req.IDs) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no IDs provided"})
		return
	}

	deletedCount, err := h.TransactionUC.DeleteTransactionsUseCase(req.IDs)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    map[string]int64{"deletedCount": deletedCount},
		"message": "Transactions deleted successfully",
	})
}
