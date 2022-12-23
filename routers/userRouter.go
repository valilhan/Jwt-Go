package routers

import (
	"net/http"
)

type UserRouter struct {
}

func NewUser() *UserRouter {
	return &UserRouter{}
}

func (h *UserRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header()
}
