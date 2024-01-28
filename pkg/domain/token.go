package domain

type UserRefreshToken struct {
	RefreshTokenID uint   `json:"refresh_token_id" gorm:"primaryKey;unique;not null"`
	RefreshToken   string `json:"refresh_token"`
	UsersID        uint
	Users          Users `gorm:"foregineKey:UsersID"`
}
