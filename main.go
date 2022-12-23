package main

import (
	"net/http"

	"github.com/gorilla/mux"
	middleware "github.com/valilhan/GolangWithJWT/middleware"
	routers "github.com/valilhan/GolangWithJWT/routers"
)

func main() {
	router := mux.NewRouter()
	router.Handle("/api-1/users/signup", routers.NewAuth()).Methods("POST")
	router.Handle("/api-1/users/login", routers.NewAuth()).Methods("POST")
	router.Handle("/api-2", routers.NewUser()).Methods("GET")
	router.Handle("/api-2", routers.NewUser()).Methods("GET")
	router.Handle("/api-2", routers.NewUser()).Methods("GET")
	router.Use(middleware.NewMiddleware().Middleware)

	http.ListenAndServe(":8080", router)
}
