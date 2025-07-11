package repository

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)


var DB *sql.DB

func ConnectDB() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file", err)
	}

	dbUrl := os.Getenv("DATABASE_URL")

	DB, err = sql.Open("postgres", dbUrl)
	if err != nil{
		log.Fatal("Error opening database:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Cannot connect to DB:", err)
	}

	fmt.Println("✅ Connected to DB successfully!")
}