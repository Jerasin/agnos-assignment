package utils

import (
	"agnos-assignment/app/config"
	"agnos-assignment/app/model"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func autoMigrate(db *gorm.DB) {
	fmt.Println("autoMigrate")
	db.AutoMigrate(&model.Patient{})
}

func InitDbClient() *gorm.DB {
	DB_HOST := config.GetEnv("DB_HOST", "localhost:3306")
	DB_NAME := config.GetEnv("DB_NAME", "api")
	DB_USER := config.GetEnv("DB_USER", "api")
	DB_PASSWORD := config.GetEnv("DB_PASSWORD", "123456")
	// DB_LOG_ENABLE := config.GetEnv("DB_LOG_ENABLE", "false")
	APP_ENV := config.GetEnv("APP_ENV", "development")

	connectionInfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable", DB_HOST, DB_USER, DB_PASSWORD, DB_NAME)
	_, err := fmt.Printf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable", DB_HOST, DB_USER, DB_PASSWORD, DB_NAME)
	// fmt.Println("n", n)
	// fmt.Println("err", err)

	if err != nil {
		panic("failed to mapping string")
	}

	dbLogLevel := logger.Info

	if APP_ENV == "development" {
		dbLogLevel = logger.Info
	} else {
		dbLogLevel = logger.Error
	}

	fmt.Println("connectionInfo", connectionInfo)
	db, err := gorm.Open(postgres.Open(connectionInfo), &gorm.Config{TranslateError: true, Logger: logger.Default.LogMode(dbLogLevel), SkipDefaultTransaction: true})
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Init Db")

	// Migrate the schema
	autoMigrate(db)

	return db
}
