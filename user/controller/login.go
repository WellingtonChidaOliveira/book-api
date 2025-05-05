package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/welligtonchida/book-api/user/models"
)

func (h *UserHandler) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		err := h.UserRepository.Login(user.Email, user.Password)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
			return
		}

		token, err := h.jwt.GenerateToken(user.ID.String(), user.Role)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
			return
		}

		c.JSON(http.StatusOK, token)
	}
}
