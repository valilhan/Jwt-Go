package controllers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
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

func HashPassword(password string) (string) {
	pass, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}
	return string(pass)
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
			return
		}
		validateErr := validate.Struct(model)
		if validateErr != nil {
			log.Println("Error with validation of model in SignUp")
			return 
		}
		countEmail, err := env.Pool.FindUserByEmail(ctx, *model.Email)
		if err != nil {
			log.Println("FindUserByEmail query error")
			return 
		}
		password := HashPassword(*model.Password)
		countPhone, err := env.Pool.FindUserByPhone(ctx, *model.Phone)
		if err != nil {
			log.Println("FindUserByPhone query error")
			return 
		}
		if countEmail > 0 || countPhone > 0 {
			log.Println("This user already exists with such email or phone")
			return 
		}
		model.Password = &password
		model.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		model.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		model.UserId = model.UserId
		token, refreshToken, err := helpers.GenerateAllTokens(*model.FirstName, *model.LastName, *model.Email, *model.UserType, model.UserId)
		if err != nil {
			log.Println("Genreting token problem")
			return
		}
		model.Token = *token
		model.RefreshToken = *refreshToken
		resultInsertionNumber, err := env.Pool.InsertUser(ctx, &model)
		if err != nil {
			log.Println("User not created")
			return
		}
		err = json.NewEncoder(w).Encode(resultInsertionNumber)
		if err != nil {
			log.Println("No error in encoding")
			return
		}
	})
}

func (env *Env) Login() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var findUser models.User
		var checkUser models.User
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		err := json.NewDecoder(r.Body).Decode(&findUser)
		if err != nil {
			log.Println("Error with decoding model in Login()")
			return
		}
		checkUser, err = env.Pool. FindUserByEmailOne(ctx, *findUser.Email)
		if err != nil {
			log.Println("Email is not correct")
			return
		}
		check, msg := VerifyPassword(*findUser.Password, *checkUser.Password)
		if check == false {
			log.Println(msg)
			return
		}
		defer cancel()
		if checkUser.Email != nil {
			 log.Println("User is not found")
			 return
		}
		token, refreshToken, err := helpers.GenerateAllTokens(*checkUser.FirstName, *checkUser.LastName, *checkUser.Email, *checkUser.UserType, checkUser.UserId)
		helpers.UpdateAllTokens(env.Pool, *token, *refreshToken, *checkUser.UserId)

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
		err := helpers.CheckUserType(r, "ADMIN"); 
		if err != nil {
			log.Println("UserType is not Admin")
			return
		}
		ctx, cancel := context.WithTimeout(context.Background(), time.Second * 10)
		recordPerPage, err := strconv.Atoi(r.Context().Value('recordPerPage'))
		if err != nil || recordPerPage < 1{
			recordPerPage = 10
		}
		page, err := strconv.Atoi(r.Context().Value('page'))
		if err != nil || page < 1 {
			page = 1
		}
		startIndex, err = strconv.Atoi(r.Context().Value('startIndex'))
		if err != nil {
			startIndex = (page - 1) * recordPerPage
		}
		defer cancel()
		users := env.Pool.SelectWithLimitOffset(ctx, startIndex, recordPerPage)
		err = json.NewEncoder(w).Encode(users)
		if err != nil {
			log.Println("Error in encoding")
			return
		}
		
	})
}
