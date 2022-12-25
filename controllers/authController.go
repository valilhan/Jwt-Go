package controllers

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	database "github.com/valilhan/GolangWithJWT/database"
	"github.com/valilhan/GolangWithJWT/helpers"
	models "github.com/valilhan/GolangWithJWT/models"
	resources "github.com/valilhan/GolangWithJWT/resources"
	"github.com/go-playground/validator/v10"
)

type ChildEnv struct {
    resources.Env
}
var validate = validator.New() 

type MyRouter mux.Router

func HashPassword(w http.ResponseWriter, r *http.Request) {

}

func VerifyPassword(w http.ResponseWriter, r *http.Request) {

}

func (env *ChildEnv) SignUp() http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request){
		ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
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
		countEmail, err := env.Pool.FindUserByEmail(ctx, model.Email)
		if err != nil {
			log.Println("FindUserByEmail query error")
		}
		
		countPhone, err := env.Pool.FindUserByPhone(ctx, model.Email)
		if err != nil {
			log.Println("FindUserByEmail query error")
		}

		if countEmail > 0 || countPhone > 0 {
			log.Println("This user already exists with such email or phone")
		}
		
	})
}

func Login(w http.ResponseWriter, r *http.Request) {

}

func (env * ChildEnv) GetUser() http.Handler {
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
func (env *main.Env) GetUsers() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		
	})
}
