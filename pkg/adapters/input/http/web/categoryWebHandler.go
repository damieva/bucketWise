package web

import (
	"bucketWise/pkg/domain"
	"bucketWise/pkg/dto"
	"bucketWise/pkg/ports"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CategoryWebHandler struct {
	CategoryUC ports.CategoryUseCase
}

func (h CategoryWebHandler) Index(c *gin.Context) {
	categories, err := h.CategoryUC.ListCategoriesUseCase("")
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	response := make([]dto.CategoryResponse, 0, len(categories))
	for _, cat := range categories {
		response = append(response, dto.CategoryResponse{
			ID:   string(cat.ID),
			Name: cat.Name,
			Type: string(cat.Type),
		})
	}

	c.HTML(http.StatusOK, "categories/index", gin.H{"Categories": response})
}

func (h CategoryWebHandler) Create(c *gin.Context) {
	cat := domain.Category{
		Name: c.PostForm("name"),
		Type: domain.CategoryType(c.PostForm("type")),
	}

	inserted, err := h.CategoryUC.CreateCategoryUseCase(cat)
	if err != nil {
		if errors.Is(err, domain.ErrCategoryAlreadyExists) {
			c.Status(http.StatusConflict)
			return
		}
		c.Status(http.StatusInternalServerError)
		return
	}

	c.HTML(http.StatusCreated, "categories/row", dto.CategoryResponse{
		ID:   string(inserted.ID),
		Name: inserted.Name,
		Type: string(inserted.Type),
	})
}

func (h CategoryWebHandler) Delete(c *gin.Context) {
	id := domain.ID(c.Param("id"))

	_, err := h.CategoryUC.DeleteCategoryUseCase([]domain.ID{id})
	if err != nil {
		if errors.Is(err, domain.ErrCategoryHasTransactions) {
			c.Status(http.StatusConflict)
			return
		}
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}
