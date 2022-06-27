package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"proletor/db/model"
)

type Database struct {
	Connection *gorm.DB
}

func Connect() *Database {
	dsn := "host=localhost user=postgres password=postgre dbname=bot-job port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	db.AutoMigrate(&model.Worker{})
	db.AutoMigrate(&model.Job{})

	if err != nil {
		panic(err)
	}

	return &Database{db}
}
