package services

import (
	"strconv"

	"github.com/welligtonchida/book-api/models"
)

type BookRepository struct {
	repo models.BookRepository
}

func NewBookRepository(repo models.BookRepository) *BookRepository {
	return &BookRepository{
		repo: repo,
	}
}

func (b *BookRepository) CreateBook(book *models.Book) error {
	// Implementation for creating a book
	return nil
}

func (b *BookRepository) GetBookByID(id string) (*models.Book, error) {
	_, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return nil, err
	}
	// Implementation for getting a book by ID
	return nil, nil
}

func (b *BookRepository) GetAllBooks() ([]models.Book, error) {
	// Implementation for getting all booksreturn []models.Book{
	return []models.Book{
		{Title: "Book 1", Author: "Author 1", Year: 2021, Description: "Description 1"},
		{Title: "Book 2", Author: "Author 2", Year: 2022, Description: "Description 2"},
	}, nil
}
func (b *BookRepository) UpdateBook(book *models.Book) error {
	// Implementation for updating a book
	return nil
}
func (b *BookRepository) DeleteBook(id string) error {
	_, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return err
	}
	// Implementation for deleting a book
	return nil
}
