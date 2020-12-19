package usecases

import (
	"context"
	"crypto/sha1"
	"fmt"
	"net/http"
	"regexp"
	"time"

	"github.com/mfaizfatah/story-tales/app/models"
	"github.com/mfaizfatah/story-tales/app/repository"
	"github.com/sirupsen/logrus"
)

const (
	// TableUser is table for user
	tableUser = "users"
)

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

// isEmailValid checks if the email provided passes the required structure and length.
func isEmailValid(e string) bool {
	if len(e) < 3 && len(e) > 254 {
		return false
	}
	return emailRegex.MatchString(e)
}

func (r *uc) Registration(ctx context.Context, req *models.User) (context.Context, *models.ResponseLogin, string, int, error) {
	var (
		sha  = sha1.New()
		res  = new(models.ResponseLogin)
		user = new(models.User)
		msg  string
		err  error
	)

	if req == nil || !isEmailValid(req.Email) {
		return ctx, nil, ErrBadRequest, http.StatusBadRequest, repository.ErrBadRequest
	}

	err = r.query.FindOne(tableUser, user, "email = ?", "id, email", req.Email)
	if user.Email != "" {
		logrus.Printf("struct user() => %v", user)
		return ctx, nil, ErrAlreadyEmail, http.StatusConflict, repository.ErrConflict
	}

	user = req

	sha.Write([]byte(user.Password))
	encrypted := sha.Sum(nil)

	user.Password = fmt.Sprintf("%x", encrypted)

	form := "2006-01-02"
	user.DateOfBirth, err = time.Parse(form, req.DateOfBirth.(string))
	if err != nil {
		return ctx, nil, ErrCreated, http.StatusInternalServerError, err
	}

	err = r.query.Insert(tableUser, user)
	if err != nil {
		return ctx, nil, ErrCreated, http.StatusInternalServerError, err
	}

	res.Message = "Registration Success"

	return ctx, res, msg, http.StatusCreated, err
}
