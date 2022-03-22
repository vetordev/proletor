package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model

	Code  string
	Price uint
}

func Connect() *gorm.DB {
	dsn := "host=localhost user=postgres password=postgre dbname=bot-job port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Error in connection")
	}

	db.AutoMigrate(&Product{})

	db.Create(&Product{Code: "D42", Price: 100})

	return db
}
