package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/truncgil/gorecta/internal/models"
	"github.com/truncgil/gorecta/pkg/database"
)

type CreatePostRequest struct {
	Title       string `json:"title" binding:"required"`
	Content     string `json:"content" binding:"required"`
	Slug        string `json:"slug" binding:"required"`
	CategoryID  uint   `json:"category_id" binding:"required"`
	TagIDs      []uint `json:"tag_ids"`
	FeaturedImg string `json:"featured_img"`
	Published   bool   `json:"published"`
}

// @Summary Create a new post
// @Description Create a new blog post with the provided details
// @Tags posts
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body CreatePostRequest true "Post creation details"
// @Success 200 {object} models.Post
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /posts [post]
func CreatePost(c *gin.Context) {
	var req CreatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, _ := c.Get("user_id")

	post := models.Post{
		Title:       req.Title,
		Content:     req.Content,
		Slug:        req.Slug,
		CategoryID:  req.CategoryID,
		UserID:      userID.(uint),
		FeaturedImg: req.FeaturedImg,
		Published:   req.Published,
	}

	db := database.GetDB()
	if err := db.Create(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create post"})
		return
	}

	// Add tags if provided
	if len(req.TagIDs) > 0 {
		var tags []models.Tag
		if err := db.Find(&tags, req.TagIDs).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find tags"})
			return
		}
		if err := db.Model(&post).Association("Tags").Replace(tags); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add tags"})
			return
		}
	}

	c.JSON(http.StatusCreated, post)
}

// @Summary Get all posts
// @Description Get a list of all blog posts
// @Tags posts
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.Post
// @Failure 500 {object} map[string]string
// @Router /posts [get]
func GetPosts(c *gin.Context) {
	var posts []models.Post
	db := database.GetDB()

	query := db.Preload("User").Preload("Category").Preload("Tags")

	// Apply filters
	if published := c.Query("published"); published != "" {
		query = query.Where("published = ?", published == "true")
	}

	if categoryID := c.Query("category_id"); categoryID != "" {
		query = query.Where("category_id = ?", categoryID)
	}

	if err := query.Find(&posts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch posts"})
		return
	}

	c.JSON(http.StatusOK, posts)
}

// @Summary Get a post by ID
// @Description Get a specific blog post by its ID
// @Tags posts
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Post ID"
// @Success 200 {object} models.Post
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /posts/{id} [get]
func GetPost(c *gin.Context) {
	id := c.Param("id")
	var post models.Post

	if err := database.GetDB().Preload("User").Preload("Category").Preload("Tags").First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	c.JSON(http.StatusOK, post)
}

// @Summary Update a post
// @Description Update an existing blog post
// @Tags posts
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Post ID"
// @Param request body CreatePostRequest true "Post update details"
// @Success 200 {object} models.Post
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /posts/{id} [put]
func UpdatePost(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	var req CreatePostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var post models.Post
	db := database.GetDB()

	if err := db.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	// Update post fields
	post.Title = req.Title
	post.Content = req.Content
	post.Slug = req.Slug
	post.CategoryID = req.CategoryID
	post.FeaturedImg = req.FeaturedImg
	post.Published = req.Published

	if err := db.Save(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update post"})
		return
	}

	// Update tags if provided
	if len(req.TagIDs) > 0 {
		var tags []models.Tag
		if err := db.Find(&tags, req.TagIDs).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find tags"})
			return
		}
		if err := db.Model(&post).Association("Tags").Replace(tags); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update tags"})
			return
		}
	}

	c.JSON(http.StatusOK, post)
}

// @Summary Delete a post
// @Description Delete a blog post by its ID
// @Tags posts
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Post ID"
// @Success 200 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /posts/{id} [delete]
func DeletePost(c *gin.Context) {
	id := c.Param("id")
	var post models.Post

	db := database.GetDB()
	if err := db.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	// Delete associated tags
	if err := db.Model(&post).Association("Tags").Clear(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete post tags"})
		return
	}

	if err := db.Delete(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete post"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully"})
}
