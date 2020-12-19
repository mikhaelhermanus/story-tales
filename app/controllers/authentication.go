package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/mfaizfatah/story-tales/app/helpers/logger"
	"github.com/mfaizfatah/story-tales/app/models"
	"github.com/mfaizfatah/story-tales/app/utils"
)

func (u *ctrl) HandlerRegistration(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Declare a new Person struct.
	var p models.User

	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		utils.Response(ctx, w, false, http.StatusBadRequest, err)
		return
	}

	ctx, res, msg, st, err := u.uc.Registration(ctx, &p)
	if err != nil {
		ctx = logger.Logf(ctx, "user registration error() => %v", err)
		utils.Response(ctx, w, false, st, msg)
		return
	}

	utils.Response(ctx, w, true, st, res)
}
