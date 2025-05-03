package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Bookhandler) DeleteBookByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseUint(c.Param("id"), 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}
		if err := h.Repo.Delete(uint(id)); err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
			return
		}
		c.JSON(http.StatusNoContent, gin.H{"message": "Book deleted successfully"})
	}
}
