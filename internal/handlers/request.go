package handlers

type CreateAccountRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateAccountResponse struct {
	Token string `json:"token"`
}

type ApiKeyRequest struct {
}

type ApiKeyResponse struct {
	ApiKey int `json:"api_key"`
}

type AverageHeightRequest struct {
}
type AverageHeightResponse struct {
	BlocksPerMinuteAVG int64 `json:"blocks_per_minute_avg"`
}
