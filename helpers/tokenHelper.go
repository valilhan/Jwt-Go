package helpers

import (
	"context"
	"log"
	"os"
	"time"
	jwt "github.com/dgrijalva/jwt-go"
	database "github.com/valilhan/GolangWithJWT/database"
)

type SignedDetail struct {
	Email string
	FirstName string
	LastName string
	UserType string
	UserId string
	jwt.StandardClaims 
}

type Env struct {
	Pool* database.PoolDB
}
var SECRET_KEY string = os.Getenv("SECRET_KEY")

func (env *Env) GenerateAllTokens(FirstName string, LastName string, Email string, UserType string, UserId string) (*string,  *string,  error) {
	claims := &SignedDetail{
		Email: Email,
		FirstName: FirstName,
		LastName: LastName,
		UserType: UserType,
		UserId: UserId ,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}
	refreshClaims := &SignedDetail{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodES256, claims).SignedString([]byte(SECRET_KEY))
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodES256, refreshClaims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		log.Panic(err)
		return nil, nil, err
	}
	return &token, &refreshToken, err
}



