package middleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/mfaizfatah/story-tales/app/adapter"
	"github.com/mfaizfatah/story-tales/app/helpers/encryption"
	"github.com/mfaizfatah/story-tales/app/helpers/logger"
)

// CheckSession middleware
func CheckSession(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		header := r.Header.Get("Authorization")

		auth := strings.Split(header, " ")

		// valueRedis, err := adapter.GetClientRedis().Get(auth[1]).Result()
		// if err != nil || err == redis.Nil {
		// 	Response(ctx, w, http.StatusUnauthorized, "Session invalid")
		// 	return
		// }

		sess := new(sessions)
		err := adapter.DBSQL().Take(sess, "session = ? AND deleted = ?", auth[1], 0).Error
		if err != nil {
			Response(ctx, w, http.StatusUnauthorized, "Session Invalid")
			return
		}

		loc, err := time.LoadLocation("Asia/Jakarta")
		if err != nil {
			Response(ctx, w, http.StatusInternalServerError, "Sorry. server configuration not available time Asia/Jakarta")
			return
		}
		now := time.Now().In(loc).Format("2006-01-02 15:04:05")

		if now > sess.ExpiredTime {
			Response(ctx, w, http.StatusUnauthorized, "Session Expired")
			return
		}

		data := strings.Split(sess.Value, "|")

		if data[0] != auth[0] {
			Response(ctx, w, http.StatusForbidden, "Bearer Token Format Error")
			return
		}

		decryptData, err := encryption.Decrypt(data[1])
		if err != nil {
			ctx = logger.Logf(ctx, "Error while decrypt header => %v", err)
			Response(ctx, w, http.StatusUnauthorized, "Wrong Authorization")
			return
		}

		ctx = logger.Logf(ctx, "authorization => %v", string(decryptData))
		ctx = logger.Device(ctx, decryptData)

		req := r.WithContext(ctx)
		req.Header.Set("x-app-token", string(decryptData))

		next.ServeHTTP(w, req)
	})
}
