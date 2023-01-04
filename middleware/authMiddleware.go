package middleware

import (
	"context"
	"log"
	"net/http"
	_ "time"

	"github.com/valilhan/GolangWithJWT/helpers"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// ctx := r.Context()
		// ctx = context.WithValue(ctx, "requestTime", time.Now().Format(time.RFC3339))
		// r = r.WithContext(ctx)
		clientToken := r.Header.Get("token")
		if clientToken == "" {
			log.Println("No token you can not access to this resources")
			return
		}
		claims, err := helpers.VerifyToken(clientToken)
		if err != "" {
			log.Println(err)
			return
		}
		ctx := r.Context()
		ctx = context.WithValue(ctx, "email", claims.Email)
		ctx = context.WithValue(ctx, "userId", claims.UserId)
		ctx = context.WithValue(ctx, "userType", claims.UserType)
		ctx = context.WithValue(ctx, "firstName", claims.FirstName)
		ctx = context.WithValue(ctx, "lastName", claims.LastName)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
