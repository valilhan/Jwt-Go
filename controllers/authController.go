package controllers

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"time"
	"github.com/gorilla/mux"
	"github.com/valilhan/GolangWithJWT/helpers"
	resources "github.com/valilhan/GolangWithJWT/resources"
	models "github.com/valilhan/GolangWithJWT/models"
)



func HashPassword(w http.ResponseWriter, r *http.Request) {

}

func VerifyPassword(w http.ResponseWriter, r *http.Request) {

}

func SignUp(w http.ResponseWriter, r *http.Request) {

}

func Login(w http.ResponseWriter, r *http.Request) {

}

func (resources.Env) GetUser() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userId := mux.Vars(r)["user_id"]
		err := helpers.MatchUserTypeToUId(r, userId)
		if err != nil {
			log.Println(err)
		}
		var user models.User
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		main.main.Env
	})
}
func (env *main.Env) GetUsers() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		
	})
}
