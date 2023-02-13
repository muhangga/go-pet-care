package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muhangga/internal/entity"
	"github.com/muhangga/internal/usecase"
	"github.com/muhangga/internal/utils"
)

type categoryDelivery struct {
	categoryUsecase usecase.CategoryUsecase
}

func NewCategoryDelivery(categoryUsecase usecase.CategoryUsecase) *categoryDelivery {
	return &categoryDelivery{
		categoryUsecase: categoryUsecase,
	}
}

func (h *categoryDelivery) CreateCategory(c *gin.Context) {
	var categoryRequest entity.CategoryRequest

	if err := c.ShouldBindJSON(&categoryRequest); err != nil {
		formatter := utils.ErrorResponse("Create category failed", http.StatusBadRequest, err.Error())

		c.JSON(http.StatusBadRequest, formatter)
		return
	}

	category, err := h.categoryUsecase.Save(categoryRequest)
	if err != nil {
		formatter := utils.ErrorResponse("Create category failed", http.StatusBadRequest, err.Error())

		c.JSON(http.StatusBadRequest, formatter)
		return
	}

	formatter := utils.SuccessResponse("Create category success", http.StatusOK, category)

	c.JSON(http.StatusOK, formatter)
}

func (h *categoryDelivery) FetchAllCategory(c *gin.Context) {
	categories, err := h.categoryUsecase.FindAll()
	if err != nil {
		formatter := utils.ErrorResponse("Failed to load category", http.StatusBadRequest, err.Error())

		c.JSON(http.StatusBadRequest, formatter)
		return
	}

	formatter := utils.SuccessResponse("Fetch All Category", http.StatusOK, categories)

	c.JSON(http.StatusOK, formatter)
}

func (h *categoryDelivery) UpdateCategory(c *gin.Context) {
	var categoryRequest entity.CategoryRequest

	if err := c.ShouldBindJSON(&categoryRequest); err != nil {
		formatter := utils.ErrorResponse("Update category failed", http.StatusBadRequest, err.Error())

		c.JSON(http.StatusBadRequest, formatter)
		return
	}

	category, err := h.categoryUsecase.Update(categoryRequest)
	if err != nil {
		formatter := utils.ErrorResponse("Update category failed", http.StatusBadRequest, err.Error())

		c.JSON(http.StatusBadRequest, formatter)
		return
	}

	formatter := utils.SuccessResponse("Update category success", http.StatusOK, category)

	c.JSON(http.StatusOK, formatter)
}

func (h *categoryDelivery) DeleteCategory(c *gin.Context) {
	id := utils.ConvertStringToInt(c.Param("id"))

	err := h.categoryUsecase.Delete(id)
	if err != nil {
		formatter := utils.ErrorResponse("Delete category failed", http.StatusBadRequest, err.Error())

		c.JSON(http.StatusBadRequest, formatter)
		return
	}

	formatter := utils.SuccessResponse("Delete category success", http.StatusOK, gin.H{"id": id})
	c.JSON(http.StatusOK, formatter)
}
