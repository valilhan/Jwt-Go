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
	Email     string
	FirstName string
	LastName  string
	UserType  string
	UserId    string
	jwt.StandardClaims
}

var SECRET_KEY string = os.Getenv("SECRET_KEY")

func GenerateAllTokens(FirstName string, LastName string, Email string, UserType string, UserId string) (*string, *string, error) {
	mySigningKey := []byte(SECRET_KEY)
	claims := &SignedDetail{
		Email:     Email,
		FirstName: FirstName,
		LastName:  LastName,
		UserType:  UserType,
		UserId:    UserId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}
	refreshClaims := &SignedDetail{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		log.Panic(err)
		return nil, nil, err
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshTokenString, err := refreshToken.SignedString(mySigningKey)
	if err != nil {
		log.Panic(err)
		return nil, nil, err
	}
	return &tokenString, &refreshTokenString, err
}
func UpdateAllTokens(db *database.PoolDB, token string, refreshToken string, UserId string) (time.Time) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	UpdatedAt, err := db.UpdateAllTokensById(ctx, token, refreshToken, UserId)
	if err != nil {
		log.Printf("Error in UpdatedAllTokensById")
	}
	return UpdatedAt
}
func VerifyToken(clientToken string) (claims *SignedDetail, msg string) {
	token, err := jwt.ParseWithClaims(clientToken, &SignedDetail{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SECRET_KEY), nil
	})
	if err != nil {
		msg = err.Error()
		return
	}
	claims, ok := token.Claims.(*SignedDetail)
	if !ok {
		msg = "The token is invalid"
		msg = msg + err.Error()
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		msg = "Token is expired"
		msg = msg + err.Error()
		return
	}
	return claims, msg
}
