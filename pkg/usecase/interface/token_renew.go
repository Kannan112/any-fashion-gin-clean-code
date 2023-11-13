package interfaces

import "context"

type RenewTokenUseCase interface {
	GetAccessToken(ctx context.Context, AccessToken string) (string, error)
	/*

	 -----> Add More function Here

	*/
}
