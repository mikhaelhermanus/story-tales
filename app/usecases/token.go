package usecases

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/mfaizfatah/story-tales/app/helpers/logger"
	"github.com/mfaizfatah/story-tales/app/models"
)

func (r *uc) GenerateToken(ctx context.Context, user *models.User) (context.Context, string, int64, error) {
	var (
		tokenValue = new(models.TokenValue)
		err        error
		token      string
		timer      = os.Getenv("TOKEN_EXPIRED")
	)

	tokenValue.IDUser = user.ID

	plaintext, err := json.Marshal(tokenValue)
	if err != nil {
		return ctx, "", 0, err
	}

	timerInt, err := strconv.Atoi(timer)
	if err != nil {
		return ctx, "", 0, err
	}

	token = uuid.New().String()
	value := base64.StdEncoding.EncodeToString(plaintext)
	exp := time.Duration(timerInt) * time.Hour

	ctx = logger.Logf(ctx, "token() => %v,%v,%v", token, value, exp)
	err = r.query.SetRedis(token, value, exp)
	if err != nil {
		return ctx, "", 0, err
	}

	duration, err := r.query.GetTTLRedis(token)
	if err != nil {
		return ctx, "", 0, err
	}

	return ctx, token, duration, nil
}

func (r *uc) GetToken(key string) (*models.TokenResponse, error) {
	var tokenReponse = new(models.TokenResponse)
	token, err := r.query.FindToken(key)
	if err != nil {
		return nil, err
	}

	duration, err := r.query.GetTTLRedis(key)
	if err != nil {
		return nil, err
	}

	tokenReponse.Key = key
	tokenReponse.Value = token
	tokenReponse.ExpiredIn = fmt.Sprintf("%v", duration)

	return tokenReponse, nil
}
