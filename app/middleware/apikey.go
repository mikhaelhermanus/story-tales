package middleware

import (
	"net/http"
	"os"
)

// CheckAPIKey middleware
func CheckAPIKey(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		apiKey := os.Getenv("API_KEY")
		header := r.Header.Get("X-API-KEY")

		if header != apiKey {
			Response(ctx, w, http.StatusUnauthorized, "Restricted Area. Please contact Administrator!")
			return
		}

		next.ServeHTTP(w, r)
		return
	})
}
