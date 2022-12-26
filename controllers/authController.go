package controllers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"
	"github.com/gorilla/mux"
	database "github.com/valilhan/GolangWithJWT/database"
	"github.com/valilhan/GolangWithJWT/helpers"
	models "github.com/valilhan/GolangWithJWT/models"
	"github.com/go-playground/validator/v10"
	"strconv"
)

type Env struct {
	Pool* database.PoolDB
}
var validate = validator.New() 

type MyRouter mux.Router

func HashPassword(w http.ResponseWriter, r *http.Request) {

}

func VerifyPassword(w http.ResponseWriter, r *http.Request) {

}

func (env *Env) SignUp() http.Handler {
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
		model.UserId = strconv.FormatInt(model.Id, 16)
		token, refreshToken, _ = helpers.GenerateAllTokens(*model.FirstName, *model.LastName, *&model.Email)

		// RefreshToken string    `json:"refreshToken"`
		// CreatedAt    time.Time `json:"createdAt"`
		// UpdatedAt    time.Time `json:"updatedAt"`
		// UserId       string    `json:"userId"`
	})
}

func Login(w http.ResponseWriter, r *http.Request) {

}

func (env * Env) GetUser() http.Handler {
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
