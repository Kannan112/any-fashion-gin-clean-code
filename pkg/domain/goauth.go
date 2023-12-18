package domain

type GoAuthUserInfo struct {
	ID         string `json:"id"`
	Email      string `json:"email"`
	Verified   bool   `json:"verified_email"`
	Name       string `json:"name"`
	GivenName  string `json:"given_name"`
	FamilyName string `json:"family_name"`
	Picture    string `json:"picture"`
	Locale     string `json:"locale"`
}
