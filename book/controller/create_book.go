package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/welligtonchida/book-api/book/models"
	"github.com/welligtonchida/book-api/repository"
)

func CreateBook(r *repository.PostgresBookRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var book models.Book
		if err := c.ShouldBindJSON(&book); err != nil {
			log.Println("Error binding JSON:", err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}
		if err := r.Create(book); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create book"})
			return
		}
		c.JSON(http.StatusCreated, book)
	}
}
