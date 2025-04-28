package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/welligtonchida/book-api/repository"
)

func GetAllBooks(s *repository.PostgresBookRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Call the repository to get all books
		books, err := s.GetAll()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch books"})
			return
		}

		// Return the books as JSON
		c.JSON(http.StatusOK, books)
	}
}
