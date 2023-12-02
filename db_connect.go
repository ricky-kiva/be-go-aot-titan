package main

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

func connectDB() (*gorm.DB, error) {
	// load environment variables file (.env)
	// errEnv := godotenv.Load(".env")

	// if errEnv != nil {
	// 	log.Fatal("Error load .env") // this will also do exit the program
	// }

	// load specific environment variable
	conn := os.Getenv("POSTGRES_URL")
	db, err := gorm.Open("postgres", conn)
	return db, err
}
