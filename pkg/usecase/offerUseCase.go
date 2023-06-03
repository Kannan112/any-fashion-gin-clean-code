package usecase

import (
	"context"

	"github.com/kannan112/go-gin-clean-arch/pkg/common/req"
)

func (c *ProductUseCase) SaveOffer(ctx context.Context, offer req.OfferTable) error {
	err := c.productRepo.SaveOffer(ctx, offer)
	// if err != nil {
	// 	return err
	// }
	return err
}
