package handlers

import (
	"bucketWise/pkg/domain"
	"bucketWise/pkg/dto"
	"bucketWise/pkg/ports"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	CategoryUC ports.CategoryUseCase
}

// CreateCategory godoc
// @Summary Create a new category
// @Description Creates a new category in the system
// @Tags categories
// @Accept json
// @Produce json
// @Param category body dto.CategoryCreateRequest true "Category data with name and type (income|expense)"
// @Success 200 {object} dto.CategoryResponse
// @Failure 400 {object} map[string]string
// @Failure 409 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /categories [post]
func (h CategoryHandler) CreateCategory(c *gin.Context) {
	var req dto.CategoryCreateRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cat := domain.Category{
		Name: req.Name,
		Type: domain.CategoryType(req.Type),
	}

	insertedCat, err := h.CategoryUC.CreateCategoryUseCase(cat)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrCategoryAlreadyExists):
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		}
		return
	}

	resp := dto.CategoryResponse{
		ID:   string(insertedCat.ID),
		Name: insertedCat.Name,
		Type: string(insertedCat.Type),
	}
	c.JSON(http.StatusCreated, resp)
}

// ListCategories godoc
// @Summary List categories (optionally filtered by name)
// @Description Retrieves all categories, or only one if the `name` query parameter is provided.
// @Tags categories
// @Param name query string false "Category name (optional)"
// @Produce json
// @Success 200 {array} dto.CategoryResponse "List of categories"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /categories [get]
func (h CategoryHandler) ListCategories(c *gin.Context) {
	name := c.Query("name")

	categoryList, err := h.CategoryUC.ListCategoriesUseCase(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var response = make([]dto.CategoryResponse, 0)
	for _, cat := range categoryList {
		response = append(response, dto.CategoryResponse{
			ID:   string(cat.ID),
			Name: cat.Name,
			Type: string(cat.Type),
		})
	}

	c.JSON(http.StatusOK, gin.H{"categories": response})
}

// DeleteCategories godoc
// @Summary Delete multiple categories
// @Description Deletes one or more categories by their IDs
// @Tags categories
// @Accept  json
// @Produce  json
// @Param body body dto.CategoriesDeleteRequest true "Array of category IDs to delete"
// @Success 200 {object} map[string]interface{} "Deletion result"
// @Failure 400 {object} map[string]string "Invalid request"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /categories [delete]
func (h CategoryHandler) DeleteCategories(c *gin.Context) {
	var req dto.CategoriesDeleteRequest

	// Validar cuerpo JSON
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	// Validar que haya al menos un ID
	if len(req.IDs) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no IDs provided"})
		return
	}

	// Convertir []string a []domain.ID
	ids := make([]domain.ID, len(req.IDs))
	for i, id := range req.IDs {
		ids[i] = domain.ID(id)
	}

	// Ejecutar caso de uso
	deletedCount, err := h.CategoryUC.DeleteCategoryUseCase(ids)
	if err != nil {

		// ID inválido
		if errors.Is(err, domain.ErrInvalidCategoryID) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Categoría no encontrada
		if errors.Is(err, domain.ErrCategoryNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		// Categoría con transacciones asociadas
		if errors.Is(err, domain.ErrCategoryHasTransactions) {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}

		// Cualquier otro error → 500
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	// Si no se borró nada (pero no hubo error → IDs válidos pero no existen)
	if deletedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "No categories found for the provided IDs",
		})
		return
	}

	// Respuesta exitosa
	c.JSON(http.StatusOK, gin.H{
		"data":    map[string]int64{"deletedCount": deletedCount},
		"message": "Categories deleted successfully",
	})
}

// UpdateCategory godoc
// @Summary Update a category
// @Description Updates the name or type of specific category
// @Tags categories
// @Accept  json
// @Produce  json
// @Param name path string true "Current category name"
// @Param category body dto.CategoryCreateRequest true "Updated category details with name and type"
// @Success 200 {object} map[string]interface{} "Update result"
// @Failure 400 {object} map[string]string "Bad request"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /categories/{name} [put]
func (h CategoryHandler) UpdateCategory(c *gin.Context) {
	catName := c.Param("name")
	var req dto.CategoryCreateRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cat := domain.Category{Name: req.Name}
	modifiedCount, err := h.CategoryUC.UpdateCategoryUseCase(catName, cat)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    map[string]int64{"modifiedCount": modifiedCount},
		"message": "Category updated successfully",
	})
}

// run pipeline
