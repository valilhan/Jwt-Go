package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/valilhan/GolangWithJWT/controllers"
	resources "github.com/valilhan/GolangWithJWT/resources"
	database "github.com/valilhan/GolangWithJWT/database"
	middleware "github.com/valilhan/GolangWithJWT/middleware"
	controllers "github.com/valilhan/GolangWithJWT/controllers"
)



func main() {
	// Initialise the connection pool as normal.
	db := database.OpenDB()
	router := mux.NewRouter()

	env := &resources.Env{database.NewPoolDB(db)}
	childEnv := &controllers.ChildEnv{*env}
	
	router.Handle("/api-1/users/signup", controllers.SignUp).Methods("POST")
	router.Handle("/api-1/users/login", controllers.Login).Methods("POST")
	router.Handle("/api-2/users", middleware.Middleware(childEnv.GetUser())).Methods("GET")
	router.Handle("/api-2/users/{user_id}", middleware.Middleware(controllers.GetUser())).Methods("GET")
	http.ListenAndServe(":8080", router)
}
