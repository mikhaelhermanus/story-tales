package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mfaizfatah/story-tales/app/helpers/logger"
)

// sessions table database
type sessions struct {
	Session     string `json:"session"`
	Value       string `json:"value"`
	ExpiredTime string `json:"expired_time"`
	Deleted     int    `json:"deleted"`
}

// response using for sending response to frontend
type response struct {
	Status       string      `json:"status"`
	ErrorMessage string      `json:"error_message"`
	Data         interface{} `json:"data"`
}

// Response ...
func Response(ctx context.Context, w http.ResponseWriter, statuscode int, data string) {
	var res response

	res.Status = "error"
	res.ErrorMessage = data
	res.Data = ""

	datares, err := json.Marshal(res)
	if err != nil {
		panic(err)
	}

	logger.EndRecord(ctx, string(datares), statuscode)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statuscode)
	fmt.Fprintf(w, string(datares))
}
