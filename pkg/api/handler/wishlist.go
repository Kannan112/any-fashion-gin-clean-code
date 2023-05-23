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

type WishlistHandler struct {
	WishlistUsecase services.WishlistUseCases
}

func NewWishlistHandler(wishlistusecase services.WishlistUseCases) *WishlistHandler {
	return &WishlistHandler{
		WishlistUsecase: wishlistusecase,
	}
}

// Wishlist
// @Summary Add Wishlist
// @ID AddToWishlist
// @Description Login as a user to access the ecommerce site
// @Tags Wishlist
// @Accept json
// @Produce json
// @Param itemId path string true "itemId"
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Router /user/wishlist/add/:itemId [post]
func (cr *WishlistHandler) AddToWishlist(c *gin.Context) {
	str := c.Param("itemId")
	itemId, err := strconv.Atoi(str)
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

	err = cr.WishlistUsecase.AddToWishlist(itemId, userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "failed to add",
			Data:       nil,
			Errors:     err.Error(),
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

// Wishlist
// @Summary Remove Item
// @ID RemoveFromWishlis
// @Description Remove item from wishlist
// @Tags Wishlist
// @Accept json
// @Produce json
// @Param itemId path string true "itemId"
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Router /user/wishlist/remove/:itemId [DELETE]

func (cr *WishlistHandler) RemoveFromWishlist(c *gin.Context) {
	str := c.Param("itemId")
	itemid, err := strconv.Atoi(str)
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
	err = cr.WishlistUsecase.RemoveFromWishlist(c, userId, itemid)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "fail to Remove",
			Data:       nil,
			Errors:     err.Error(),
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

func (c *WishlistHandler) ListAllWishlist(ctx *gin.Context) {
	count, err1 := strconv.Atoi(ctx.Query("count"))
	page, err2 := strconv.Atoi(ctx.Query("page"))
	if err1 != nil || err2 != nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "page not found",
			Data:       nil,
			Errors:     err1,
		})
		return
	}
	var pagenation req.Pagenation
	pagenation.Count = count
	pagenation.Page = page

	userId, err := handlerUtil.GetUserIdFromContext(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "Cant find the user",
			Data:       nil,
			Errors:     err,
		})
		return
	}
	wishlist, err := c.WishlistUsecase.ListAllWishlist(ctx, userId,pagenation)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "fail to list",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusBadRequest, res.Response{
		StatusCode: 200,
		Message:    "Wishlist",
		Data:       wishlist,
		Errors:     nil,
	})
}
