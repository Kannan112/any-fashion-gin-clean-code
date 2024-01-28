package res

type Token struct {
	Access_token  string
	Refresh_token string
}

type TokenCalim struct {
	ID   uint
	Role string
}

type AdminToken struct {
	Token string
}
