package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/welligtonchida/book-api/book/models"
	"github.com/welligtonchida/book-api/repository"
)

type Bookhandler struct {
	Repo repository.BookRepository
}

func NewBookHandler(repo repository.BookRepository) *Bookhandler {
	return &Bookhandler{
		Repo: repo,
	}
}

func (h *Bookhandler) CreateBook() gin.HandlerFunc {
	return func(c *gin.Context) {
		var book models.Book
		if err := c.ShouldBindJSON(&book); err != nil {
			log.Println("Error binding JSON:", err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}
		if err := h.Repo.Create(book); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create book"})
			return
		}
		c.JSON(http.StatusCreated, book)
	}
}
