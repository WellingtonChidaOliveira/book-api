package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/welligtonchida/book-api/services"
)

func DeleteBookByID(s *services.BookRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		if err := s.DeleteBook(id); err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
	}
}
