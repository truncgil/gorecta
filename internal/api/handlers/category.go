package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/truncgil/gorecta/internal/models"
	"github.com/truncgil/gorecta/pkg/database"
)

type CreateCategoryRequest struct {
	Name        string `json:"name" binding:"required"`
	Slug        string `json:"slug" binding:"required"`
	Description string `json:"description"`
}

// @Summary Create a new category
// @Description Create a new category with the provided details
// @Tags categories
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body CreateCategoryRequest true "Category creation details"
// @Success 200 {object} models.Category
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /categories [post]
func CreateCategory(c *gin.Context) {
	var req CreateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category := models.Category{
		Name:        req.Name,
		Slug:        req.Slug,
		Description: req.Description,
	}

	if err := database.GetDB().Create(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create category"})
		return
	}

	c.JSON(http.StatusCreated, category)
}

// @Summary Get all categories
// @Description Get a list of all categories
// @Tags categories
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.Category
// @Failure 500 {object} map[string]string
// @Router /categories [get]
func GetCategories(c *gin.Context) {
	var categories []models.Category
	if err := database.GetDB().Find(&categories).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch categories"})
		return
	}

	c.JSON(http.StatusOK, categories)
}

// @Summary Get a category by ID
// @Description Get a specific category by its ID
// @Tags categories
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Category ID"
// @Success 200 {object} models.Category
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /categories/{id} [get]
func GetCategory(c *gin.Context) {
	id := c.Param("id")
	var category models.Category

	if err := database.GetDB().First(&category, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	c.JSON(http.StatusOK, category)
}

// @Summary Update a category
// @Description Update an existing category
// @Tags categories
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Category ID"
// @Param request body CreateCategoryRequest true "Category update details"
// @Success 200 {object} models.Category
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /categories/{id} [put]
func UpdateCategory(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}

	var req CreateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var category models.Category
	if err := database.GetDB().First(&category, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	category.Name = req.Name
	category.Slug = req.Slug
	category.Description = req.Description

	if err := database.GetDB().Save(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update category"})
		return
	}

	c.JSON(http.StatusOK, category)
}

// @Summary Delete a category
// @Description Delete a category by its ID
// @Tags categories
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Category ID"
// @Success 200 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /categories/{id} [delete]
func DeleteCategory(c *gin.Context) {
	id := c.Param("id")
	var category models.Category

	if err := database.GetDB().First(&category, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	if err := database.GetDB().Delete(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete category"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category deleted successfully"})
}
