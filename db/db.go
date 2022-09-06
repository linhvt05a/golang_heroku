package db

import (
	"example/kenva-be/models"
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	// connection string
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("err loading: %v", err)
	}
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	user := os.Getenv("POST_GRE_USER")
	password := os.Getenv("PASSWORD")
	dbname := os.Getenv("POST_GRE_DATABASE")

	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	fmt.Print(psqlconn)
	db, err := gorm.Open(postgres.Open(psqlconn), &gorm.Config{})
	fmt.Println("Connected!")
	db.AutoMigrate(&models.User{})

	return db
}
