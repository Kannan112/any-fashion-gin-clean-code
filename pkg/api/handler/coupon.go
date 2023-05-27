package handler

import (
	"fmt"
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

// CreateCoupon
// @Summary Admin can create new coupon
// @ID create-coupon
// @Description Admin can create new coupons
// @Tags Coupon
// @Accept json
// @Produce json
// @Param new_coupon_details body req.Coupons true "details of new coupon to be created"
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Router /admin/coupon/add [post]
func (cr *CouponHandler) AddCoupon(c *gin.Context) {

	var newCoupon req.Coupons
	err := c.ShouldBindJSON(&newCoupon)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "failed to bind coupon data",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	fmt.Println("Add Coupon", newCoupon.DiscountMaximumAmount)
	err = cr.CouponUseCase.AddCoupon(c, newCoupon)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "failed to create coupon",
			Data:       nil,
			Errors:     err.Error(),
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

// DeleteCoupon
// @Summary Delete a coupon
// @ID DeleteCoupon
// @Description Delete coupon
// @Tags Coupon
// @Accept json
// @Produce json
// @Param couponId path string true "New Admin details"
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Router /admin/coupon/delete/:couponId [delete]
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
			Message:    "failed to delete coupon",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusBadRequest, res.Response{
		StatusCode: 200,
		Message:    "Coupon deleted",
		Data:       nil,
		Errors:     nil,
	})
}

// UpdateCoupon
// @Summary Update a existing coupon
// @ID UpdateCoupon
// @Description admin coupon update
// @Tags Coupon
// @Accept json
// @Produce json
// @Param admin body req.Coupons true "New Admin details"
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Router /admin/coupon/update [patch]
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
	var coupon req.Coupons
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
			StatusCode: 400,
			Message:    "failed to update",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, res.Response{
		StatusCode: 200,
		Message:    "Coupon updated",
		Data:       nil,
		Errors:     err,
	})

}

// ViewCoupon
// @Summary view coupons
// @ID ViewCoupon
// @Description admin view all coupon
// @Tags Coupon
// @Accept json
// @Produce json
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Router /admin/coupon [get]
func (c *CouponHandler) ViewCoupon(ctx *gin.Context) {
	coupon, err := c.CouponUseCase.ViewCoupon(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "failed to display",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusAccepted, res.Response{
		StatusCode: 200,
		Message:    "List of coupons",
		Data:       coupon,
		Errors:     nil,
	})

}
