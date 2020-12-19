package middleware

import (
	"net/http"
	"time"

	"github.com/mfaizfatah/story-tales/app/helpers/logger"
)

// RecordMiddleware is using for record any request from outside
func RecordMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		loc, err := time.LoadLocation("Asia/Jakarta")
		if err != nil {
			panic(err)
		}

		start := time.Now().In(loc)

		req := logger.StartRecord(r, start)

		next.ServeHTTP(w, req)
	})
}
