package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	handlerUtil "github.com/kannan112/go-gin-clean-arch/pkg/api/handlerUril"
	"github.com/kannan112/go-gin-clean-arch/pkg/common/res"
	services "github.com/kannan112/go-gin-clean-arch/pkg/usecase/interface"
)

type WishlistHandler struct {
	WishlistUsecase services.WishlistUseCases
}

func NewWishlistHandler(wishlistusecase services.WishlistUseCases) *WishlistHandler {
	return &WishlistHandler{
		WishlistUsecase: wishlistusecase,
	}
}
func (cr *WishlistHandler) AddToWishlist(c *gin.Context) {
	str := c.Param("productId")
	productid, err := strconv.Atoi(str)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "cant get the product id",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	userId, err := handlerUtil.GetUserIdFromContext(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "Cant find the user",
			Data:       nil,
			Errors:     err,
		})
		return
	}

	err = cr.WishlistUsecase.AddToWishlist(productid, userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "failed to add",
			Data:       nil,
			Errors:     err,
		})
		return
	}
	c.JSON(http.StatusAccepted, res.Response{
		StatusCode: 200,
		Message:    "added to wishlist",
		Data:       nil,
		Errors:     nil,
	})
}

func (cr *WishlistHandler) RemoveFromWishlist(c *gin.Context) {
	str := c.Param("productId")
	productid, err := strconv.Atoi(str)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "cant get the product id",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	userId, err := handlerUtil.GetUserIdFromContext(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "Cant find the user",
			Data:       nil,
			Errors:     err,
		})
		return
	}
	err = cr.WishlistUsecase.RemoveFromWishlist(c, userId, productid)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "fail to Remove",
			Data:       nil,
			Errors:     err,
		})
		return
	}
	c.JSON(http.StatusAccepted, res.Response{
		StatusCode: 200,
		Message:    "Remove from wishlist",
		Data:       nil,
		Errors:     nil,
	})
}
