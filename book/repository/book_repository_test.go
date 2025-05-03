package repository

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"github.com/welligtonchida/book-api/book/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupTestPostgresDB(t *testing.T) (*gorm.DB, func()) {
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

	postgresC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		t.Fatalf("failed to start postgres container: %s", err)
	}

	host, err := postgresC.Host(ctx)
	if err != nil {
		t.Fatalf("failed to get postgres container host: %s", err)
	}
	port, err := postgresC.MappedPort(ctx, "5432")
	if err != nil {
		t.Fatalf("failed to get postgres container port: %s", err)
	}

	dsn := "host=" + host + " user=postgres password=postgres dbname=bookapi_test port=" + port.Port() + " sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect to postgres: %s", err)
	}
	if err := db.AutoMigrate(&models.Book{}); err != nil {
		t.Fatalf("failed to migrate database: %s", err)
	}

	teardown := func() {
		postgresC.Terminate(ctx)
	}

	return db, teardown

}

func TestPostgresBookRepository(t *testing.T) {
	db, teardown := setupTestPostgresDB(t)
	defer teardown()

	repo := NewBookRepository(db)

	book := models.Book{
		Title:       "The Go Programming Language",
		Author:      "Alan A. A. Donovan",
		Year:        2015,
		Description: "An introduction to the Go programming language.",
	}

	book.ID = 1 // Reset ID to 1 to ensure a new record is created

	// Test Create
	err := repo.Create(book)
	assert.NoError(t, err)
	assert.NotZero(t, book.ID)

	// Test GetByID
	retrievedBook, err := repo.GetByID(book.ID)
	assert.NoError(t, err)
	assert.Equal(t, book.Title, retrievedBook.Title)
	assert.Equal(t, book.Author, retrievedBook.Author)
	assert.Equal(t, book.Year, retrievedBook.Year)
	assert.Equal(t, book.Description, retrievedBook.Description)

	// Test GetAll
	books, err := repo.GetAll()
	assert.NoError(t, err)
	assert.NotEmpty(t, books)
	assert.Greater(t, len(books), 0)

	// Test Update
	book.Title = "The Go Programming Language (Updated)"
	err = repo.Update(book)
	assert.NoError(t, err)
	retrievedBook, err = repo.GetByID(book.ID)
	assert.NoError(t, err)
	assert.Equal(t, book.Title, retrievedBook.Title)

	// Test Delete
	err = repo.Delete(book.ID)
	assert.NoError(t, err)
	retrievedBook, err = repo.GetByID(book.ID)
	assert.Error(t, err)
	assert.Equal(t, "book not found", err.Error())
	assert.Equal(t, models.Book{}, retrievedBook)
}

func TestPostgresRespositoryErrors(t *testing.T) {
	db, teardown := setupTestPostgresDB(t)
	defer teardown()

	repo := NewBookRepository(db)

	book := models.Book{
		Title:       "The Go Programming Language",
		Author:      "Alan A. A. Donovan",
		Year:        2015,
		Description: "An introduction to the Go programming language.",
	}
	book.ID = 1 // Reset ID to 1 to ensure a new record is created

	// Test Create
	err := repo.Create(book)
	assert.NoError(t, err)
	assert.NotZero(t, book.ID)

	// Test GetByID with non-existing ID
	_, err = repo.GetByID(2)
	assert.Error(t, err)
	assert.Equal(t, "book not found", err.Error())
	// Test Update with non-existing ID
	book.ID = 2
	err = repo.Update(book)
	assert.Error(t, err)
	assert.Equal(t, "book not found", err.Error())
	// Test Delete with non-existing ID
	err = repo.Delete(2)
	assert.Error(t, err)
	assert.Equal(t, "book not found", err.Error())
}
