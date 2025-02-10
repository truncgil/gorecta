package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/truncgil/gorecta/internal/api/handlers"
	"github.com/truncgil/gorecta/internal/api/middleware"
)

// SetupRoutes configures all the routes for our application
func SetupRoutes(router *gin.Engine) {
	// API v1 group
	v1 := router.Group("/api/v1")

	// Auth routes
	auth := v1.Group("/auth")
	{
		auth.POST("/register", handlers.Register)
		auth.POST("/login", handlers.Login)
	}

	// Protected routes
	protected := v1.Group("")
	protected.Use(middleware.AuthMiddleware())
	{
		// Posts routes
		posts := protected.Group("/posts")
		{
			posts.GET("", handlers.GetPosts)
			posts.POST("", middleware.RoleMiddleware("admin", "editor"), handlers.CreatePost)
			posts.GET("/:id", handlers.GetPost)
			posts.PUT("/:id", middleware.RoleMiddleware("admin", "editor"), handlers.UpdatePost)
			posts.DELETE("/:id", middleware.RoleMiddleware("admin"), handlers.DeletePost)
		}

		// Categories routes
		categories := protected.Group("/categories")
		{
			categories.GET("", handlers.GetCategories)
			categories.POST("", middleware.RoleMiddleware("admin"), handlers.CreateCategory)
			categories.GET("/:id", handlers.GetCategory)
			categories.PUT("/:id", middleware.RoleMiddleware("admin"), handlers.UpdateCategory)
			categories.DELETE("/:id", middleware.RoleMiddleware("admin"), handlers.DeleteCategory)
		}

		// Tags routes (to be implemented)
		tags := protected.Group("/tags")
		{
			tags.GET("")
			tags.POST("", middleware.RoleMiddleware("admin"))
			tags.GET("/:id")
			tags.PUT("/:id", middleware.RoleMiddleware("admin"))
			tags.DELETE("/:id", middleware.RoleMiddleware("admin"))
		}

		// Users routes (to be implemented)
		users := protected.Group("/users")
		{
			users.GET("", middleware.RoleMiddleware("admin"))
			users.GET("/:id")
			users.PUT("/:id")
			users.DELETE("/:id", middleware.RoleMiddleware("admin"))
		}
	}
}
