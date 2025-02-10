package handlers

import (
	"net/http"
	"strconv"

	"gorecta-cms/internal/models"
	"gorecta-cms/pkg/database"

	"github.com/gin-gonic/gin"
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

// CreatePost creates a new post
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

// GetPosts returns a list of posts
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

// GetPost returns a single post by ID
func GetPost(c *gin.Context) {
	id := c.Param("id")
	var post models.Post

	if err := database.GetDB().Preload("User").Preload("Category").Preload("Tags").First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	c.JSON(http.StatusOK, post)
}

// UpdatePost updates a post
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

// DeletePost deletes a post
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
