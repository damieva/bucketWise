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
// @Param category body dto.CategoryCreateRequest true "Category data"
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

	cat := domain.Category{Name: req.Name}
	insertedId, err := h.CategoryUC.CreateCategoryUseCase(cat)
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
		ID:   insertedId.(string),
		Name: req.Name,
	}
	c.JSON(http.StatusOK, resp)
}

// ListAllCategories godoc
// @Summary      List all categories
// @Description  Retrieves all categories stored in the system.
// @Tags         categories
// @Produce      json
// @Success      200  {object}  map[string][]dto.CategoryResponse  "List of categories"
// @Failure      500  {object}  map[string]string  "Internal server error"
// @Router       /categories [get]
func (h CategoryHandler) ListAllCategories(c *gin.Context) {
	categoryList, err := h.CategoryUC.ListAllCategoryUseCase()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"categories": categoryList})
}

// GetCategoryByName godoc
// @Summary Get category by name
// @Description Retrieves the details of a specific category by its name
// @Tags categories
// @Param name path string true "Category name"
// @Produce  json
// @Success 200 {object} map[string]interface{} "Category details"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /categories/{name} [get]
func (h CategoryHandler) GetCategoryByName(c *gin.Context) {
	name := c.Param("name")
	catName := domain.Category{Name: name}

	result, err := h.CategoryUC.ListOneCategoryUseCase(catName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"category": result})
}

// DeleteCategory godoc
// @Summary Delete a category
// @Description Deletes a specific category by its name
// @Tags categories
// @Param name path string true "Category name"
// @Produce  json
// @Success 200 {object} map[string]interface{} "Number of deleted categories"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /categories/{name} [delete]
func (h CategoryHandler) DeleteCategory(c *gin.Context) {
	name := c.Param("name")
	catName := domain.Category{Name: name}

	deletedCount, err := h.CategoryUC.DeleteCategoryUseCase(catName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Amount of categories deleted": deletedCount})
}

// UpdateCategory godoc
// @Summary Update a category
// @Description Updates the name or details of a specific category
// @Tags categories
// @Accept  json
// @Produce  json
// @Param name path string true "Current category name"
// @Param category body dto.CategoryCreateRequest true "Updated category details"
// @Success 200 {object} map[string]interface{} "Number of modified categories"
// @Failure 400 {object} map[string]string "Bad request"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /categories/{name} [put]
func (h CategoryHandler) UpdateCategory(c *gin.Context) {
	catName := c.Param("name")
	var categoryUpdateParms domain.Category
	if err := c.BindJSON(&categoryUpdateParms); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	modifiedCount, err := h.CategoryUC.UpdateCategoryUseCase(catName, categoryUpdateParms)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Amount of categories modified": modifiedCount})
}
