package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/valilhan/GolangWithJWT/controllers"
	resources "github.com/valilhan/GolangWithJWT/resources"
	database "github.com/valilhan/GolangWithJWT/database"
	middleware "github.com/valilhan/GolangWithJWT/middleware"
)



func main() {
	// Initialise the connection pool as normal.
	db := database.OpenDB()
	router := mux.NewRouter()

	env := &resources.Env{database.NewPoolDB(db)}
	
	router.Handle("/api-1/users/signup", controllers.SignUp).Methods("POST")
	router.Handle("/api-1/users/login", controllers.Login).Methods("POST")
	router.Handle("/api-2/users", middleware.Middleware(controllers.GetUsers())).Methods("GET")
	router.Handle("/api-2/users/{user_id}", middleware.Middleware(controllers.GetUser())).Methods("GET")
	http.ListenAndServe(":8080", router)
}
