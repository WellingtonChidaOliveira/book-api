package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/welligtonchida/book-api/services"
)

func GetBookByID(s *services.BookRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		book, err := s.GetBookByID(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
			return
		}
		c.JSON(http.StatusOK, book)
	}
}
