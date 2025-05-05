package config

import (
	bm "github.com/welligtonchida/book-api/book/models"
	um "github.com/welligtonchida/book-api/user/models"
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

func MigrateDatabase(db *gorm.DB) error {
	if err := db.AutoMigrate(&bm.Book{}); err != nil {
		return err
	}

	if err := db.AutoMigrate(&um.User{}); err != nil {
		return err
	}

	return nil
}

func StartDatabase() (*gorm.DB, error) {
	db, err := InitDatabase()
	if err != nil {
		return nil, err
	}

	db.Exec("CREATE EXTENSION IF NOT EXISTS 'uuid-ossp';")

	if err = MigrateDatabase(db); err != nil {
		return nil, err
	}

	return db, nil
}
