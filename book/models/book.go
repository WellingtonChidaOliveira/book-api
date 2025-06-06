package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title       string `json:"title" gorm:"type:varchar(100);not null"`
	Author      string `json:"author" gorm:"type:varchar(100);not null"`
	Year        int    `json:"year" gorm:"type:int;not null"`
	Description string `json:"description" gorm:"type:text;not null"`
}
