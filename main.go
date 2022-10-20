package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Did string `gorm:"typevarchar(100);unique_index"`
}

type StreamIndex struct {
	gorm.Model

	StreamId string `gorm:"typevarchar(100);unique_index"`
	Did string
}

var db *gorm.DB
var err error

func main() {
	godotenv.Load()
	// Loading environment variables
	host := os.Getenv("HOST")
	dbPort := os.Getenv("DBPORT")
	user := os.Getenv("USER")
	dbName := os.Getenv("NAME")
	password := os.Getenv("PASSWORD")
	// fmt.Println("env: ",host,dbPort, user,dbName,password)

	// Database connection string
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", host, user, password, dbName, dbPort)
	// fmt.Println("DSN: ", dsn)

	// Openning connection to database
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	} else {
		fmt.Println("Successfully connected to database!")
	}

	// Migrate the schema
	db.AutoMigrate(&Account{})
	db.AutoMigrate(&StreamIndex{})
}