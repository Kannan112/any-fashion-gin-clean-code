package interfaces

import "context"

type RefreshTokenRepository interface {
	AdminRefreshTokenAdd(token string, adminID uint) error
	AdminFindRefreshToken(ctx context.Context, adminID uint) (string, error)

	UserRefreshTokenAdd(token string, userID uint) error
	UserFindRefreshToken(ctx context.Context, userID uint) (string, error)
}
