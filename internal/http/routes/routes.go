package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/welligtonchida/book-api/book/controller"
	br "github.com/welligtonchida/book-api/book/repository"
	"github.com/welligtonchida/book-api/config"
)

func Handlers() *gin.Engine {
	// Define your routes hera
	r := gin.Default()
	d, err := config.InitDatabase()
	if err != nil {
		panic(err)
	}
	repo := br.NewBookRepository(d)
	if err != nil {
		panic(err)
	}

	bookHandler := controller.NewBookHandler(repo)
	r.GET("/api/v1/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})
	r.GET("api/v1/books", bookHandler.GetAllBooks())
	r.GET("api/v1/books/:id", bookHandler.GetBookByID())
	r.POST("api/v1/books", bookHandler.CreateBook())
	r.DELETE("api/v1/books/:id", bookHandler.DeleteBookByID())
	r.PUT("api/v1/books/:id", bookHandler.UpdateBook())
	return r
}
