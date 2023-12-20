package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	handlerUtil "github.com/kannan112/go-gin-clean-arch/pkg/api/handlerUril"
	"github.com/kannan112/go-gin-clean-arch/pkg/common/req"
	"github.com/kannan112/go-gin-clean-arch/pkg/common/res"
	services "github.com/kannan112/go-gin-clean-arch/pkg/usecase/interface"
)

type CartHandler struct {
	cartUsecase services.CartUseCases
}

func NewCartHandler(cartUsecases services.CartUseCases) *CartHandler {
	return &CartHandler{
		cartUsecase: cartUsecases,
	}
}

// AddToCart
// @Summary User add product-item to cart
// @ID add-to-cart
// @Description User can add product item to the cart
// @Tags Cart
// @Accept json
// @Produce json
// @Param product_items_id path string true "product_items_id"
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Router /api/user/cart/add/{product_items_id} [post]
func (cr *CartHandler) AddToCart(c *gin.Context) {
	userId, err := handlerUtil.GetUserIdFromContext(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "cant find userid",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	paramsid := c.Param("product_items_id")
	productId, err := strconv.Atoi(paramsid)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "cant find",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	err = cr.cartUsecase.AddToCart(productId, userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "cant add product into cart",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{
		"Added": "successfully",
	})
}

// remove from cart

// RemoveFromCart
// @Summary user remove item
// @ID remove-from-cart
// @Description User can remove product item from carts
// @Tags Cart
// @Accept json
// @Produce json
// @Param product_item_id path string true "product_item_id"
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Router /api/user/cart/remove/{product_item_id} [delete]
func (cr *CartHandler) RemoveFromCart(c *gin.Context) {
	userId, err := handlerUtil.GetUserIdFromContext(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "cant find userid",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	paramsId := c.Param("product_item_id")
	productitemid, err := strconv.Atoi(paramsId)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "cant find productid",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	err = cr.cartUsecase.RemoveFromCart(userId, productitemid)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "cant remove from cart",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, res.Response{
		StatusCode: 200,
		Message:    "item removed from cart",
		Data:       nil,
		Errors:     nil,
	})
}

// Llist cart

// List Cart
// @Summary user can view
// @ID list-cart
// @Description User can view the cart with amount
// @Tags Cart
// @Accept json
// @Produce json
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Router /api/user/cart/list [get]
func (cr *CartHandler) ListCart(c *gin.Context) {
	userId, err := handlerUtil.GetUserIdFromContext(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "failed to get userId",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	list, err := cr.cartUsecase.ListCart(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "Failed",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusAccepted, res.Response{
		StatusCode: 200,
		Message:    "List Cart",
		Data:       list,
		Errors:     nil,
	})
}

// List Cart items
// @Summary user can view
// @ID list-cart-items
// @Description User can view items in cart
// @Tags Cart
// @Accept json
// @Produce json
// @Param count query int false "Page number for pagination"
// @Param page query int false "Number of items to retrieve per page"
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Router /user/cart-item/list [get]
func (c *CartHandler) ListCartItems(ctx *gin.Context) {
	var pagenation req.Pagenation
	countStr := ctx.Query("count")
	pageStr := ctx.Query("page")
	if countStr != "" || pageStr != "" {
		count, err1 := strconv.Atoi(countStr)
		page, err := strconv.Atoi(pageStr)
		pagenation.Count = count
		pagenation.Page = page
		if err != nil || err1 != nil {
			ctx.JSON(http.StatusBadRequest, res.Response{
				StatusCode: 400,
				Message:    "page not found",
				Data:       nil,
				Errors:     err,
			})
			return
		}
	}
	userId, err := handlerUtil.GetUserIdFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "failed to get userId",
			Data:       nil,
		})
		return
	}
	list, err := c.cartUsecase.ListCartItems(ctx, userId, pagenation)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "Failed",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, res.Response{
		StatusCode: 200,
		Message:    "cart product details",
		Data:       list,
		Errors:     nil,
	})
}
