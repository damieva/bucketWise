package handlers

import (
	"bucketWise/pkg/domain"
	"bucketWise/pkg/ports"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	CategoryUC ports.CategoryUseCase
}

func (h CategoryHandler) CreateCategory(c *gin.Context) {
	// El handler ha de encargarse de:
	//   - Traducir el request
	//   - Validación (casos de uso)
	//   - Consumir el servicio
	//   - Traducir el response
	var categoryCreateParms domain.Category
	if err := c.BindJSON(&categoryCreateParms); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	// ========================

	// Definimos arriba una vble de tipo interfaz ports.CreateCategoryUseCase y aquí usamos el metodo create
	insertedId, err := h.CategoryUC.CreateCategoryUseCase(categoryCreateParms)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "oops!"})
	}

	// =======================
	// Envía una respuesta HTTP con el código 200 y el category_id
	c.JSON(http.StatusOK, gin.H{"category_id": insertedId})
}

func (h CategoryHandler) ListAllCategories(c *gin.Context) {
	categoryList, err := h.CategoryUC.ListAllCategoryUseCase()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "oops!"})
	}

	// =======================
	// Envía una respuesta HTTP con el código 200 y el category_id
	c.JSON(http.StatusOK, gin.H{"categories": categoryList})
}
