package repository

import (
	"errors"

	"github.com/welligtonchida/book-api/book/models"
	"gorm.io/gorm"
)

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{
		db: db,
	}

}

func (r *BookRepository) Create(book models.Book) error {
	return r.db.Create(&book).Error
}

func (r *BookRepository) GetByID(id uint) (models.Book, error) {
	var book models.Book
	err := r.db.First(&book, "id = ?", id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return book, errors.New("book not found")
	}
	return book, err
}

func (r *BookRepository) GetAll() ([]models.Book, error) {
	var books []models.Book
	err := r.db.Find(&books).Error
	return books, err
}

func (r *BookRepository) Update(book models.Book) error {
	_, err := r.GetByID(book.ID)
	if err != nil {
		return errors.New("book not found")
	}
	return r.db.Save(&book).Error
}

func (r *BookRepository) Delete(id uint) error {
	_, err := r.GetByID(id)
	if err != nil {
		return errors.New("book not found")
	}
	return r.db.Delete(&models.Book{}, "id = ?", id).Error
}
