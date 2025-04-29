package repository

import (
	"errors"

	"github.com/welligtonchida/book-api/book/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type BookRepository interface {
	Create(book models.Book) error
	GetByID(id uint) (models.Book, error)
	GetAll() ([]models.Book, error)
	Update(book models.Book) error
	Delete(id uint) error
}

type PostgresBookRepository struct {
	db *gorm.DB
}

func InitDatabase() (*gorm.DB, error) {
	conn := "host=localhost user=postgres password=postgres dbname=bookapi port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func NewPostgresRepository(db *gorm.DB) (*PostgresBookRepository, error) {
	err := db.AutoMigrate(&models.Book{})
	if err != nil {
		return nil, err
	}

	return &PostgresBookRepository{
		db: db,
	}, nil
}

func (r *PostgresBookRepository) Create(book models.Book) error {
	return r.db.Create(&book).Error
}

func (r *PostgresBookRepository) GetByID(id uint) (models.Book, error) {
	var book models.Book
	err := r.db.First(&book, "id = ?", id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return book, errors.New("book not found")
	}
	return book, err
}

func (r *PostgresBookRepository) GetAll() ([]models.Book, error) {
	var books []models.Book
	err := r.db.Find(&books).Error
	return books, err
}

func (r *PostgresBookRepository) Update(book models.Book) error {
	_, err := r.GetByID(book.ID)
	if err != nil {
		return errors.New("book not found")
	}
	return r.db.Save(&book).Error
}

func (r *PostgresBookRepository) Delete(id uint) error {
	_, err := r.GetByID(id)
	if err != nil {
		return errors.New("book not found")
	}
	return r.db.Delete(&models.Book{}, "id = ?", id).Error
}
