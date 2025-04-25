package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/welligtonchida/book-api/controller"
	"github.com/welligtonchida/book-api/services"
)

func Handlers() *gin.Engine {
	// Define your routes hera
	r := gin.Default()
	s := services.NewBookRepository(nil)
	r.GET("/api/v1/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})
	r.GET("api/v1/books", controller.GetAllBooks(s))
	r.GET("api/v1/books/:id", controller.GetBookByID(s))
	r.POST("api/v1/books", controller.CreateBook(s))
	r.DELETE("api/v1/books/:id", controller.DeleteBookByID(s))
	r.PUT("api/v1/books/:id", controller.UpdateBook(s))
	return r
}
