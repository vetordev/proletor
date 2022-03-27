package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	Connection *gorm.DB
}

func Connect() *Database {
	dsn := "host=localhost user=postgres password=postgre dbname=bot-job port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Error in connection")
	}

	return &Database{db}
}
