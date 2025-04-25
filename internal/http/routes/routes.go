package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/welligtonchida/book-api/controller"
)

func Handlers() *gin.Engine {
	// Define your routes hera
	r := gin.Default()
	r.GET("/api/v1/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	r.GET("/api/v1/book", controller.GetBook)
	r.Handle("GET", "/api/v1/book/:id", controller.GetBook)
	return r
}
