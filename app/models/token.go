package models

// TokenValue is model data for token value
type TokenValue struct {
	IDUser int `json:"id_user"`
}

// TokenResponse is model data for response token
type TokenResponse struct {
	Key       string `json:"key"`
	Value     string `json:"value"`
	ExpiredIn string `json:"expired_in"`
}
