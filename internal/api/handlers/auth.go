package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/gorecta-cms/internal/models"
	"github.com/yourusername/gorecta-cms/pkg/auth"
	"github.com/yourusername/gorecta-cms/pkg/database"
)

type RegisterRequest struct {
	Name     string `