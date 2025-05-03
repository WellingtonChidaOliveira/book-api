package repository

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

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
