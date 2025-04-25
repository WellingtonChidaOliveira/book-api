package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/welligtonchida/book-api/models"
	"github.com/welligtonchida/book-api/services"
)

func CreateBook(s *services.BookRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var book models.Book
		if err := c.ShouldBindJSON(&book); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}
		if err := s.CreateBook(&book); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create book"})
			return
		}
		c.JSON(http.StatusCreated, book)
	}
}
