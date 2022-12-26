package main

import (
	"net/http"
	"github.com/gorilla/mux"
	database "github.com/valilhan/GolangWithJWT/database"
	middleware "github.com/valilhan/GolangWithJWT/middleware"
	controllers "github.com/valilhan/GolangWithJWT/controllers"
)



func main() {
	// Initialise the connection pool as normal.
	db := database.OpenDB()
	router := mux.NewRouter()

	env := &controllers.Env{database.NewPoolDB(db)}
	
	router.Handle("/api-1/users/signup", env.SignUp()).Methods("POST")
	router.Handle("/api-1/users/login", env.Login).Methods("POST")
	router.Handle("/api-2/users", middleware.Middleware(env.GetUsers())).Methods("GET")
	router.Handle("/api-2/users/{user_id}", middleware.Middleware(env.GetUser())).Methods("GET")
	http.ListenAndServe(":8080", router)
}
