package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Bookhandler) GetAllBooks() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Call the repository to get all books
		books, err := h.Repo.GetAll()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch books"})
			return
		}

		// Return the books as JSON
		c.JSON(http.StatusOK, books)
	}
}
