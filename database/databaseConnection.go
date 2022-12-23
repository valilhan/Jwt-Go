package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"
	"github.com/joho/godotenv"
)
var (
    DB_USER     = os.Getenv("DB_USER")
    DB_PASSWORD = os.Getenv("DB_PASSWORD")
    DB_NAME     = os.Getenv("DB_NAME")
    DB_HOST     = os.Getenv("DB_HOST")
    DB_PORT     = os.Getenv("DB_PORT")
)
func openDB() *sql.Conn {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		log.Fatal("Error with opening databalse")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Millisecond)
	conn, err := db.Conn(ctx)
	if err != nil {
		log.Fatal("Error with connecting database ")
	}
	defer cancel()
	log.Println("Database successfully connected")
	return conn
}
