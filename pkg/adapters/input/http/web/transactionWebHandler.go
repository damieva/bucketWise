package web

import (
	"bucketWise/pkg/dto"
	"bucketWise/pkg/ports"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransactionWebHandler struct {
	TransactionUC ports.TransactionUseCase
}

func (h TransactionWebHandler) Index(c *gin.Context) {
	transactions, err := h.TransactionUC.ListTransactionsUseCase("")
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	response := make([]dto.TransactionResponse, 0, len(transactions))
	for _, tx := range transactions {
		response = append(response, dto.TransactionResponse{
			ID:           string(tx.ID),
			Amount:       tx.Amount,
			Description:  tx.Description,
			Date:         tx.Date,
			CategoryID:   tx.CategoryID,
			CategoryName: tx.CategoryName,
			Type:         string(tx.Type),
		})
	}

	c.HTML(http.StatusOK, "transactions/index", gin.H{"Transactions": response})
}
