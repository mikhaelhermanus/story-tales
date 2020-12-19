package middleware

import (
	"net/http"

	"github.com/mfaizfatah/story-tales/app/helpers/logger"
)

// NotFound ...
func NotFound(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	logger.EndRecord(ctx, "Route Not Found", http.StatusNotFound)

	w.Write([]byte("Not found"))
}

// NotAllowed ...
func NotAllowed(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	logger.EndRecord(ctx, "Method Not Allowed", http.StatusMethodNotAllowed)

	w.Write([]byte("Method Not Allowed"))
}
