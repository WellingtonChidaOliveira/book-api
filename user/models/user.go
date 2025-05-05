package models

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	Name     string    `json:"name" gorm:"type:varchar(100);not null"`
	Email    string    `json:"email" gorm:"type:varchar(100);unique;not null"`
	Password string    `json:"password" gorm:"type:varchar(100);not null"`
	Role     string    `json:"role" gorm:"type:varchar(50);not null"`
}

type UserRepository interface {
	Create(user *User) (*User, error)
	GetByID(id string) (*User, error)
	GetByEmail(email string) (*User, error)
	Update(user *User) (*User, error)
	Delete(id string) error
	GetAll() ([]User, error)
	GetByRole(role string) ([]User, error)
	Login(email, password string) error
}
