package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/welligtonchida/book-api/book/models"
)

func (h *Bookhandler) UpdateBook() gin.HandlerFunc {
	return func(c *gin.Context) {
		//id := c.Param("id")
		var book models.Book
		if err := c.ShouldBindJSON(&book); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		err := h.Repo.Update(book)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update book"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Book updated successfully"})
	}
}
