package repository

import (
	"github.com/welligtonchida/book-api/user/models"
	"gorm.io/gorm"
)

type PostgresUserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) models.UserRepository {
	return &PostgresUserRepository{
		db: db,
	}
}

func (r *PostgresUserRepository) Create(user *models.User) (*models.User, error) {
	if err := r.db.Create(user).Error; err != nil {
		return &models.User{}, err
	}

	return user, nil
}

func (r *PostgresUserRepository) Login(email, password string) error {
	var user models.User
	if err := r.db.Where("email = ? AND password = ?", email, password).First(&user).Error; err != nil {
		return err
	}

	return nil
}

func (r *PostgresUserRepository) GetByID(id string) (*models.User, error) {
	var user models.User
	if err := r.db.First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *PostgresUserRepository) GetByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.db.First(&user, "email = ?", email).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *PostgresUserRepository) Update(user *models.User) (*models.User, error) {
	if err := r.db.Save(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *PostgresUserRepository) Delete(id string) error {
	if err := r.db.Delete(&models.User{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *PostgresUserRepository) GetAll() ([]models.User, error) {
	var users []models.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *PostgresUserRepository) GetByRole(role string) ([]models.User, error) {
	var users []models.User
	if err := r.db.Where("role = ?", role).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
