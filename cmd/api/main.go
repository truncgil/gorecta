package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/yourusername/gorecta-cms/internal/api/routes"
	"github.com/yourusername/gorecta-cms/internal/models"
	"github.com/yourusername/gorecta-cms/pkg/database"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found")
	}
}

func main() {
	// Set Gin mode
	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Initialize database
	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Auto migrate database schemas
	if err := db.AutoMigrate(
		&models.User{},
		&models.Post{},
		&models.Category{},
		&models.Tag{},
	); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Initialize router
	router := gin.Default()

	// CORS configuration
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", os.Getenv("ALLOWED_ORIGINS"))
		c.Writer.Header().Set("Access-Control-Allow-Methods", os.Getenv("ALLOWED_METHODS"))
		c.Writer.Header().Set("Access-Control-Allow-Headers", os.Getenv("ALLOWED_HEADERS"))

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// Setup routes
	routes.SetupRoutes(router)

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	// Get server address from environment variables
	serverAddr := fmt.Sprintf("%s:%s",
		os.Getenv("SERVER_HOST"),
		os.Getenv("SERVER_PORT"),
	)

	// Start server
	log.Printf("Server starting on %s", serverAddr)
	if err := router.Run(serverAddr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
