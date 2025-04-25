package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBook(c *gin.Context) {
	book := map[string]string{
		"title":  "The Great Gatsby",
		"author": "F. Scott Fitzgerald",
	}

	// Return the book as JSON
	c.JSON(http.StatusOK, book)

}
