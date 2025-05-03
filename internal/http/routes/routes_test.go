package routes

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"github.com/welligtonchida/book-api/book/controller"
	"github.com/welligtonchida/book-api/book/models"
	"github.com/welligtonchida/book-api/repository"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setUpDataBase(t *testing.T) (*gorm.DB, func()) {
	ctx := context.Background()

	req := testcontainers.ContainerRequest{
		Image:        "postgres:latest",
		ExposedPorts: []string{"5432/tcp"},
		Env: map[string]string{
			"POSTGRES_USER":     "postgres",
			"POSTGRES_PASSWORD": "postgres",
			"POSTGRES_DB":       "bookapi_test",
		},
		WaitingFor: wait.ForListeningPort("5432/tcp").WithStartupTimeout(30 * time.Second),
	}

	// Start the container
	c, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})

	if err != nil {
		t.Fatalf("Failed to start container: %s", err)
	}
	host, _ := c.Host(ctx)
	port, _ := c.MappedPort(ctx, "5432")

	// Create a connection string
	dsn := "host=" + host + " port=" + port.Port() + " user=postgres password=postgres dbname=bookapi_test sslmode=disable"
	// Connect to the database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to database: %s", err)
	}
	// Run migrations or setup your database schema here
	db.AutoMigrate(&models.Book{})

	return db, func() {
		c.Terminate(ctx)
	}
}

func setUpRouter(repo *repository.PostgresBookRepository) *gin.Engine {
	r := gin.Default()
	r.GET("/api/v1/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	bookHandler := controller.NewBookHandler(repo)

	r.GET("api/v1/books", bookHandler.GetAllBooks())
	r.GET("api/v1/books/:id", bookHandler.GetBookByID())
	r.POST("api/v1/books", bookHandler.CreateBook())
	r.DELETE("api/v1/books/:id", bookHandler.DeleteBookByID())
	r.PUT("api/v1/books/:id", bookHandler.UpdateBook())
	return r
}

func TestHandlers(t *testing.T) {
	db, teardown := setUpDataBase(t)
	defer teardown()

	repo, err := repository.NewPostgresRepository(db)
	assert.NoError(t, err)

	r := setUpRouter(repo)

	newBook := models.Book{
		Title:       "Clean Architecture",
		Author:      "Robert C. Martin",
		Year:        2017,
		Description: "Software architecture is a set of structures needed to reason about the system.",
	}
	newBook.ID = 1 // Ensure ID is zero for new book creation
	body, _ := json.Marshal(newBook)

	// Test Create Book
	req := httptest.NewRequest("POST", "/api/v1/books", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)

	var createdBook models.Book
	err = json.Unmarshal(w.Body.Bytes(), &createdBook)
	assert.NoError(t, err)
	assert.Equal(t, newBook.Title, createdBook.Title)
	assert.NotZero(t, createdBook.ID)

	// Test Get All Books
	req = httptest.NewRequest("GET", "/api/v1/books", nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	var books []models.Book
	err = json.Unmarshal(w.Body.Bytes(), &books)
	assert.NoError(t, err)
	assert.Greater(t, len(books), 0)

	// Test Get Book By ID
	req = httptest.NewRequest("GET", fmt.Sprintf("/api/v1/books/%d", createdBook.ID), nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	var book models.Book
	err = json.Unmarshal(w.Body.Bytes(), &book)
	assert.NoError(t, err)
	assert.Equal(t, createdBook.ID, book.ID)
	// Test Update Book
	updatedBook := models.Book{
		Title:       "Clean Code",
		Author:      "Robert C. Martin",
		Year:        2008,
		Description: "A Handbook of Agile Software Craftsmanship",
	}
	updatedBook.ID = createdBook.ID // Use the same ID for update
	body, _ = json.Marshal(updatedBook)
	req = httptest.NewRequest("PUT", fmt.Sprintf("/api/v1/books/%d", createdBook.ID), bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	// Test Delete Book
	req = httptest.NewRequest("DELETE", fmt.Sprintf("/api/v1/books/%d", createdBook.ID), nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNoContent, w.Code)
	// Verify the book is deleted
	req = httptest.NewRequest("GET", fmt.Sprintf("/api/v1/books/%d", createdBook.ID), nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
	var notFoundResponse map[string]string
	err = json.Unmarshal(w.Body.Bytes(), &notFoundResponse)
	assert.NoError(t, err)
	assert.Equal(t, "Book not found", notFoundResponse["error"])

	// Test Health Check
	req = httptest.NewRequest("GET", "/api/v1/health", nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	var healthResponse map[string]string
	err = json.Unmarshal(w.Body.Bytes(), &healthResponse)
	assert.NoError(t, err)
	assert.Equal(t, "ok", healthResponse["status"])
}
