package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/valilhan/GolangWithJWT/models"
)

// init is invoked before main()

var DB_HOST, DB_NAME, DB_PASSWORD, DB_USER string
var DB_PORT int

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
	DB_USER = os.Getenv("DB_USER")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	DB_NAME = os.Getenv("DB_NAME")
	DB_HOST = os.Getenv("DB_HOST")
	DB_PORT, _ = strconv.Atoi(os.Getenv("DB_PORT"))
}

type PoolDB struct {
	db *sqlx.DB
}

func NewPoolDB(inDB *sqlx.DB) *PoolDB {
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
func (pool *PoolDB) SelectWithLimitOffset(ctx context.Context, startIndex int, recordPerPage int) ([]models.User, error) {
	var users []models.User
	query := `SELECT * FROM USERS LIMIT $1 OFFSET $2`
	rows, err := pool.db.QueryxContext(ctx, query, recordPerPage, startIndex)
	if err != nil {
		log.Println(err)
		log.Println("Error in execute query with SelectWithLimitOffset")
	}
	for rows.Next() {
		var user models.User
		err = rows.StructScan(&user)
		if err != nil {
			log.Println(err)
			log.Println("Error in rows scanning in SelectWithLimitOffset")
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
func (pool *PoolDB) UpdateAllTokensById(ctx context.Context, token string, refreshToken string, UserId string) (time.Time, error) {
	UpdatedAt, err := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	if err != nil {
		log.Println("Getting time UpdatedAt with error")
	}
	query := `UPDATE USERS SET token = $1, refreshtoken = $2, updatedat = $3 WHERE userId = $4`
	// UPDATE table_name
	// SET column1 = value1, column2 = value2, ...
	// WHERE condition;
	_, err = pool.db.ExecContext(ctx, query, token, refreshToken, UpdatedAt, UserId)
	if err != nil {
		log.Println(err)
		log.Println("Error with updating refreshtoken and token")
	}
	return UpdatedAt, err

}
func (pool *PoolDB) InsertUser(ctx context.Context, user *models.User) (int, error) {
	var lastId int
	query := `INSERT INTO USERS (id, firstname, lastname, password, email, phone,  usertype, token, refreshtoken, createdat, updatedat, userid) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) RETURNING userid`
	err := pool.db.QueryRowContext(ctx, query, user.Id, user.FirstName, user.LastName, user.Password, user.Email, user.Phone, user.UserType, user.Token, user.RefreshToken, user.CreatedAt, user.UpdatedAt, user.UserId).Scan(&lastId)
	if err != nil {
		log.Println("InsertUser query error")
		return -1, err
	}
	return lastId, nil
}
func (pool *PoolDB) FindUserByEmailOne(ctx context.Context, email string) (models.User, error) {
	var user models.User
	query := `SELECT * FROM USERS WHERE email = $1;`
	err := pool.db.QueryRowxContext(ctx, query, email).StructScan(&user)
	if err != nil {
		log.Println("Error in FindUserByEmailOne")
		log.Println(err)
	}
	return user, nil
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
	query := "SELECT * FROM USERS WHERE userId = $1"
	var model models.User
	err := pool.db.QueryRowxContext(ctx, query, user_id).StructScan(&model)
	if err != nil {
		log.Println("GetUser query error")
		return nil, err
	}
	return &model, nil
}

func OpenDB() *sqlx.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sqlx.Open("postgres", psqlconn)
	if err != nil {
		log.Fatal("Error with opening databalse")
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("Can not establish connection with database")
	}
	log.Println("Database connection established")
	return db
}
