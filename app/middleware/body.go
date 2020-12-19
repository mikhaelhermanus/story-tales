package middleware

import (
	"net/http"
	"reflect"

	"github.com/mfaizfatah/story-tales/app/helpers/encryption"
	"github.com/mfaizfatah/story-tales/app/helpers/logger"
)

// CheckBody middleware
func CheckBody(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		data := r.FormValue("data")

		if reflect.ValueOf(data).IsZero() {
			next.ServeHTTP(w, r)
			return
		}

		decrypt, err := encryption.Decrypt(data)
		if err != nil {
			ctx = logger.Logf(ctx, "Error while decrypt body => %v", err)
			Response(ctx, w, http.StatusBadRequest, "Decryption failed!")
			return
		}

		ctx = logger.Logf(ctx, "Decrypt body => %s", string(decrypt))

		req := r.WithContext(ctx)
		req.Form.Set("data", string(decrypt))

		next.ServeHTTP(w, req)
	})
}
