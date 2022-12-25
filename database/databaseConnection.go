package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"github.com/valilhan/GolangWithJWT/models"
	"github.com/joho/godotenv"
)

var (
	DB_USER     = os.Getenv("DB_USER")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	DB_NAME     = os.Getenv("DB_NAME")
	DB_HOST     = os.Getenv("DB_HOST")
	DB_PORT     = os.Getenv("DB_PORT")
)

type PoolDB struct {
	db *sql.DB
}

func NewPoolDB(inDB *sql.DB) *PoolDB {
	return &PoolDB{
		db: inDB,
	}
}
func (pool *PoolDB) FindUserByPhone(ctx context.Context, phone string) (int, error) {
	query := `SELECT COUNT(userId) FROM USERS WHERE phone = $1;`
	var count int
	err := pool.db.QueryRowContext(ctx, query, phone).Scan(&count)
	if err != nil {
		log.Println("FindUserByEmail query error")
		return -1, err
	}
	return count, nil
}
func (pool *PoolDB) FindUserByEmail(ctx context.Context, email string) (int, error) {
	query := `SELECT COUNT(userId) FROM USERS WHERE email = $1;`
	var count int
	err := pool.db.QueryRowContext(ctx, query, email).Scan(&count)
	if err != nil {
		log.Println("FindUserByEmail query error")
		return -1, err
	}
	return count, nil
}

func (pool *PoolDB) GetUser(ctx context.Context, user_id string) (*models.User, error) {
	query := "SELECT * FROM USERS WHERE userId = $1 RETURNING"
	var model models.User
	err := pool.db.QueryRowContext(ctx, query, user_id).Scan(&model)
	if err != nil {
		log.Println("GetUser query error")
		return nil, err
	}
	return &model, nil
} 

func OpenDB() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		log.Fatal("Error with opening databalse")
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("Can not establish connection with database")
	}
	log.Println("Database connection established")
	return db
}
