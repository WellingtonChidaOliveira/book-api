package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/welligtonchida/book-api/internal/middleware/jwt"
	"github.com/welligtonchida/book-api/user/models"
)

type UserHandler struct {
	UserRepository models.UserRepository
	jwt            *jwt.JWT
}

func NewUserHandler(userRepository models.UserRepository, jwt *jwt.JWT) *UserHandler {
	return &UserHandler{
		UserRepository: userRepository,
		jwt:            jwt,
	}
}

func (h *UserHandler) CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		createdUser, err := h.UserRepository.Create(&user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
			return
		}

		// Generate JWT token
		token, err := h.jwt.GenerateToken(createdUser.ID.String(), createdUser.Role)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
			return
		}

		c.JSON(http.StatusCreated, token)
	}
}
