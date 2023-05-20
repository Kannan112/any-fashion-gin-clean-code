package handler

import (
	"net/http"
	"strconv"

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
func (c *CouponHandler) DeleteCoupon(ctx *gin.Context) {
	strId := ctx.Param("couponId")
	couponId, err := strconv.Atoi(strId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "fail to get id",
			Data:       nil,
			Errors:     err,
		})
		return
	}
	err = c.CouponUseCase.DeleteCoupon(ctx, couponId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "failed to create coupon",
			Data:       nil,
			Errors:     err,
		})
		return
	}
	ctx.JSON(http.StatusBadRequest, res.Response{
		StatusCode: 200,
		Message:    "Coupon deleted",
		Data:       nil,
		Errors:     err,
	})
}
func (c *CouponHandler) UpdateCoupon(ctx *gin.Context) {
	strId := ctx.Param("couponId")
	couponId, err := strconv.Atoi(strId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "fail to get id",
			Data:       nil,
			Errors:     err,
		})
		return
	}
	var coupon req.Coupon
	err = ctx.Bind(&coupon)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "failed to bind coupon data",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	err = c.CouponUseCase.UpdateCoupon(ctx, coupon, couponId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 200,
			Message:    "Coupon Updated",
			Data:       nil,
			Errors:     nil,
		})
	}

}
