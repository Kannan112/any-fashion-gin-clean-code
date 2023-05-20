package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kannan112/go-gin-clean-arch/pkg/common/req"
	"github.com/kannan112/go-gin-clean-arch/pkg/common/res"
	services "github.com/kannan112/go-gin-clean-arch/pkg/usecase/interface"
)

type CouponHandler struct {
	CouponUseCase services.CouponUseCase
}

func NewCouponHandler(CouponUsecase services.CouponUseCase) *CouponHandler {
	return &CouponHandler{
		CouponUseCase: CouponUsecase,
	}
}
func (cr *CouponHandler) AddCoupon(c *gin.Context) {
	var coupon req.Coupon
	err := c.Bind(&coupon)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "failed to bind coupon data",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	err = cr.CouponUseCase.AddCoupon(c, coupon)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "failed to create coupon",
			Data:       nil,
			Errors:     err,
		})
		return
	}
	c.JSON(http.StatusAccepted, res.Response{
		StatusCode: 200,
		Message:    "created coupon",
		Data:       nil,
		Errors:     err,
	})
}
