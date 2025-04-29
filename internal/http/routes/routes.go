package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/welligtonchida/book-api/book/controller"
	"github.com/welligtonchida/book-api/repository"
)

func Handlers() *gin.Engine {
	// Define your routes hera
	r := gin.Default()
	d, err := repository.InitDatabase()
	if err != nil {
		panic(err)
	}
	repo, err := repository.NewPostgresRepository(d)
	if err != nil {
		panic(err)
	}
	r.GET("/api/v1/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})
	r.GET("api/v1/books", controller.GetAllBooks(repo))
	r.GET("api/v1/books/:id", controller.GetBookByID(repo))
	r.POST("api/v1/books", controller.CreateBook(repo))
	r.DELETE("api/v1/books/:id", controller.DeleteBookByID(repo))
	r.PUT("api/v1/books/:id", controller.UpdateBook(repo))
	return r
}
