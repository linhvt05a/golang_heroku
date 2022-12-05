package db

import (
	"example/kenva-be/configs"
	"example/kenva-be/models"
	"fmt"
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	config, errEnv := configs.LoadConfig(".")
	if errEnv != nil {
		log.Fatalf("could not load config: %v", errEnv)
	}
	db, err := gorm.Open(postgres.Open(config.PostgresUrl), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect DB")
		return nil
	}
	fmt.Println("*********------------Database is Connected!---------------*******")
	db.AutoMigrate(&models.User{})
	fmt.Println("*********------------AutoMigrate Models!---------------*******")
	return db
}
