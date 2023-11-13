package req

type AccessToken struct {
	TokenString string `json:"token_string" binding:"required"`
}
