package controllers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	database "github.com/valilhan/GolangWithJWT/database"
	helpers "github.com/valilhan/GolangWithJWT/helpers"
	models "github.com/valilhan/GolangWithJWT/models"
	"golang.org/x/crypto/bcrypt"
)

type Env struct {
	Pool *database.PoolDB
}

var validate = validator.New()

type MyRouter mux.Router

func HashPassword(w http.ResponseWriter, r *http.Request) {

}

func VerifyPassword(originalPassword string, checkPassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(originalPassword), []byte(checkPassword))
	if err != nil {
		msg := "Password is not correct"
		return false, msg
	}
	return true, ""
}

func (env *Env) SignUp() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		var model models.User
		err := json.NewDecoder(r.Body).Decode(&model)
		if err != nil {
			log.Println("Error with decoding model in SignUp")
		}
		validateErr := validate.Struct(model)
		if validateErr != nil {
			log.Println("Error with validation of model in SignUp")
		}
		countEmail, err := env.Pool.FindUserByEmail(ctx, *model.Email)
		if err != nil {
			log.Println("FindUserByEmail query error")
		}

		countPhone, err := env.Pool.FindUserByPhone(ctx, *model.Email)
		if err != nil {
			log.Println("FindUserByEmail query error")
		}

		if countEmail > 0 || countPhone > 0 {
			log.Println("This user already exists with such email or phone")
		}
		model.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		model.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		model.UserId = model.UserId
		env := helpers.Env{env.Pool}
		token, refreshToken, _ := env.GenerateAllTokens(*model.FirstName, *model.LastName, *model.Email, *model.UserType, model.UserId)
		model.Token = *token
		model.RefreshToken = *refreshToken

		resultInsertionNumber, err := env.Pool.InsertUser(ctx, &model)
		if err != nil {
			log.Println("User not created")
		}
		err = json.NewEncoder(w).Encode(resultInsertionNumber)
		if err != nil {
			log.Println("No error in encoding")
		}
	})
}

func (env *Env) Login() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var findUser models.User
		var checkUser *models.User
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		err := json.NewDecoder(r.Body).Decode(&findUser)
		if err != nil {
			log.Println("Error with decoding model in Login()")
		}
		checkUser, err = env.Pool.FindUserByEmailOne(ctx, *findUser.Email)
		if err != nil {
			log.Println("Email is not correct")
		}
		check, msg := VerifyPassword(*findUser.Password, *checkUser.Password)
		if check == false {
			log.Println(msg)
		}
		defer cancel()
	})
}

func (env *Env) GetUser() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userId := mux.Vars(r)["user_id"]
		err := helpers.MatchUserTypeToUId(r, userId)
		if err != nil {
			log.Println(err)
		}
		var user *models.User
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		user, err = env.Pool.GetUser(ctx, userId)
		if err != nil {
			log.Println(err, "GetUser query errors")
		}
		json.NewEncoder(w).Encode(&user)
	})
}
func (env *Env) GetUsers() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}
