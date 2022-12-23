package middleware

import "net/http"

type authMiddleware struct {

}

func (auth *authMiddleware) ServeHTTP(w http.ResponseWriter, r * http.Request) {

}

func NewMiddleware() *authMiddleware{
	return &authMiddleware{}
}

func (auth *authMiddleware) Middleware(next http.Handler) http.Handler {
	return &authMiddleware{}
}