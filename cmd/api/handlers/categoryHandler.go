package handlers

import (
	"bucketWise/pkg/domain"
	"bucketWise/pkg/ports"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	CategoryUC ports.CategoryUseCase
}

func (h CategoryHandler) CreateCategory(c *gin.Context) {
	var categoryCreateParms domain.Category
	if err := c.BindJSON(&categoryCreateParms); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	insertedId, err := h.CategoryUC.CreateCategoryUseCase(categoryCreateParms)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrCategoryAlreadyExists):
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"category_id": insertedId})
}

func (h CategoryHandler) ListAllCategories(c *gin.Context) {
	categoryList, err := h.CategoryUC.ListAllCategoryUseCase()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "oops!"})
	}

	c.JSON(http.StatusOK, gin.H{"categories": categoryList})
}

func (h CategoryHandler) GetCategoryByName(c *gin.Context) {
	name := c.Param("name")
	category := domain.Category{Name: name}

	result, err := h.CategoryUC.ListOneCategoryUseCase(category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "oops!"})
	}

	c.JSON(http.StatusOK, gin.H{"category": result})
}

func (h CategoryHandler) DeleteCategory(c *gin.Context) {
	name := c.Param("name")
	category := domain.Category{Name: name}

	deletedCount, err := h.CategoryUC.DeleteCategoryUseCase(category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "oops!"})
	}

	c.JSON(http.StatusOK, gin.H{"Amount of categories deleted": deletedCount})
}

/*
func (h CategoryHandler) UpdateCategory(c *gin.Context) {
	var categoryUpdateParms domain.Category
	if err := c.BindJSON(&categoryUpdateParms); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	insertedId, err := h.CategoryUC.UpdateCategoryUseCase(categoryUpdateParms)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "oops!"})
	}

	c.JSON(http.StatusOK, gin.H{"category_id": insertedId})
}
*/
