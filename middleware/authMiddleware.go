package middleware

import (
	"net/http"
	_ "context"
	_ "time"
)
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		// ctx := r.Context()
		// ctx = context.WithValue(ctx, "requestTime", time.Now().Format(time.RFC3339))
        // r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}