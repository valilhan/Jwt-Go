package routers

import (
	"net/http"
)

// Handler
type AuthRouter struct {
}

func NewAuth() *AuthRouter {
	return &AuthRouter{}
}

// ServeHTTP determines that AuthRouter is Handler
func (h *AuthRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
