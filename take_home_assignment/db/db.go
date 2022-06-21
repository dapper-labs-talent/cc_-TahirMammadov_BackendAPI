package db

import (
	"github.com/go-playground/log/v8"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/tmammado/take-home-assignment/model"
)

func Init() *gorm.DB {
	log.Info("Initializing DB")

	url := "postgres://postgres:password@localhost:5432/dapper_labs"

	db, err := gorm.Open(postgres.Open(url), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&model.User{})

	return db
}
