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

type ProductHandler struct {
	productuseCase services.ProductUseCase
}

func NewProductHandler(productUseCase services.ProductUseCase) *ProductHandler {
	return &ProductHandler{
		productuseCase: productUseCase,
	}
}

// CreateCategory
// @Summary Create new product category
// @ID create-category
// @Description Admin can create new category from admin panel
// @Tags Product Category
// @Accept json
// @Produce json
// @Param category_name body req.Category true "New category name"
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Router /admin/category/add [post]

func (cr *ProductHandler) CreateCategory(c *gin.Context) {
	var category req.Category
	err := c.Bind(&category)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "bind faild",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	NewCategory, err := cr.productuseCase.CreateCategory(category)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "Can't creat category",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, res.Response{
		StatusCode: 200,
		Message:    "Category Created",
		Data:       NewCategory,
		Errors:     nil,
	})

}

// UpdateCategory
// @Summary Admin can update category details
// @ID update-category
// @Description Admin can update category details
// @Tags Product Category
// @Accept json
// @Produce json
// @Param id path string true "ID of the Category to be updated"
// @Param category_details body req.Category true "category info"
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Router /admin/category/update/{id} [patch]

func (cr *ProductHandler) UpdatCategory(c *gin.Context) {
	var category req.Category
	err := c.Bind(&category)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "bind faild",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	paramsId := c.Param("id")
	id, err := strconv.Atoi(paramsId)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "failed to get",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	updateCategory, err := cr.productuseCase.UpdateCategory(category, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "Can't update category",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, res.Response{
		StatusCode: 200,
		Message:    "Category Updated",
		Data:       updateCategory,
		Errors:     nil,
	})
}

// DeleteCategory
// @Summary Admin can delete a category
// @ID delete-category
// @Description Admin can delete a category
// @Tags Product Category
// @Accept json
// @Produce json
// @Param category_id path string true "category_id"
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Router /admin/category/delete/{category_id} [delete]
func (cr *ProductHandler) DeleteCategory(c *gin.Context) {
	paramID := c.Param("category_id")
	id, err := strconv.Atoi(paramID)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "cant bind data",
			Data:       nil,
			Errors:     err.Error(),
		})
	}
	err = cr.productuseCase.DeleteCategory(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "cant delete category",
			Data:       nil,
			Errors:     err.Error(),
		})
	}

	c.JSON(http.StatusOK, res.Response{
		StatusCode: 200,
		Message:    "Category deleted",
		Data:       nil,
		Errors:     nil,
	})

}

// ListAllCategories
// @Summary View all available categories
// @ID view-all-categories
// @Description Admin, users and unregistered users can see all the available categories
// @Tags Product Category
// @Accept json
// @Produce json
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Router /admin/category/listall [get]

func (cr *ProductHandler) ListCategories(c *gin.Context) {

	count, err1 := strconv.Atoi(c.Query("count"))
	page, err2 := strconv.Atoi(c.Query("page"))
	if err1 != nil || err2 != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "cant bind data",
			Data:       nil,
			Errors:     err1,
		})
		return
	}
	var pagenation req.Pagenation
	pagenation.Count = count
	pagenation.Page = page
	fmt.Printf("count: %v,page: %v", count, page)
	categories, err := cr.productuseCase.ListCategories(c, pagenation)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "can't find category",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, res.Response{
		StatusCode: 200,
		Message:    "Ctegories are",
		Data:       categories,
		Errors:     nil,
	})
}

// FindCategoryByID
// @Summary Fetch details of a specific category using category id
// @ID find-category-by-id
// @Description Users and admins can fetch details of a specific category using id
// @Tags Product Category
// @Accept json
// @Produce json
// @Param id path string true "category id"
// @Success 200 {object} res.Response
// @Failure 422 {object} res.Response
// @Failure 500 {object} res.Response
// @Router /admin/category/find/{id} [get]

func (cr *ProductHandler) DisplayCategory(c *gin.Context) {
	var category []res.Product
	paramsId := c.Param("category_id")
	id, err := strconv.Atoi(paramsId)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "can't find id",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	category, err = cr.productuseCase.DisplayCategory(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "can't find category",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, res.Response{
		StatusCode: 200,
		Message:    "Category is",
		Data:       category,
		Errors:     nil,
	})

}

// --------ProductItem--------

// CreateProductItem
// @Summary Creates a new product item
// @ID create-product-item
// @Description This endpoint allows an admin user to create a new product item.
// @Tags Product Item
// @Accept json
// @Produce json
// @Param product_item body req.ProductItem true "Product item details"
// @Success 200 {object} res.Response "Successfully added new product item"
// @Failure 400 {object} res.Response "Failed to add new product item"
// @Router /admin/product-item/add/ [post]

func (cr *ProductHandler) AddProduct(c *gin.Context) {
	var product req.Product
	err := c.Bind(&product)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "Cant bind",
			Data:       nil,
			Errors:     err,
		})
		return
	}

	newProduct, err := cr.productuseCase.AddProduct(product)

	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "Cant create",
			Data:       nil,
			Errors:     err,
		})
		return

	}
	c.JSON(http.StatusOK, res.Response{
		StatusCode: 200,
		Message:    "product created",
		Data:       newProduct,
		Errors:     nil,
	})
}
func (cr *ProductHandler) UpdateProduct(c *gin.Context) {
	var product req.Product
	productId := c.Param("id")
	id, err := strconv.Atoi(productId)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "cant get the id",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	err = c.Bind(&product)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "binding failed",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	productUpdate, err := cr.productuseCase.UpdateProduct(id, product)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "failed to get the productDetails",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusAccepted, res.Response{
		StatusCode: 200,
		Message:    "Updated successfully",
		Data:       productUpdate,
		Errors:     nil,
	})
}
func (cr *ProductHandler) DeleteProduct(c *gin.Context) {
	productId := c.Param("id")
	id, err := strconv.Atoi(productId)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "cant get the id",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	err = cr.productuseCase.DeleteProduct(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "failed to get the productDetails",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusAccepted, res.Response{
		StatusCode: 200,
		Message:    "Deleted sucessfuly",
		Data:       nil,
		Errors:     nil,
	})
}
func (cr *ProductHandler) ListProducts(c *gin.Context) {
	product, err := cr.productuseCase.ListProducts()
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "failed to get the productDetails",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusAccepted, res.Response{
		StatusCode: 200,
		Message:    "List of product",
		Data:       product,
		Errors:     nil,
	})
}

func (cr *ProductHandler) DisplayProduct(c *gin.Context) {
	productId := c.Param("id")
	id, err := strconv.Atoi(productId)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "failed get productid",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	product, err := cr.productuseCase.DisplayProduct(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "failed",
			Data:       nil,
			Errors:     err.Error(),
		})
	}
	c.JSON(http.StatusAccepted, res.Response{
		StatusCode: 200,
		Message:    "success",
		Data:       product,
		Errors:     nil,
	})
}

func (cr *ProductHandler) AddProductItem(c *gin.Context) {
	var productItem req.ProductItem
	err := c.Bind(&productItem)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "Cant bind",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	newProductItem, err := cr.productuseCase.AddProductItem(productItem)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "Cant create",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, res.Response{
		StatusCode: 200,
		Message:    "product created",
		Data:       newProductItem,
		Errors:     nil,
	})
}
func (cr *ProductHandler) UpdateProductItem(c *gin.Context) {
	var productItem req.ProductItem
	err := c.Bind(&productItem)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "Cant bind",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	paramsId := c.Param("id")
	id, err := strconv.Atoi(paramsId)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "can't find id",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	UpdateProductItem, err := cr.productuseCase.UpdateProductItem(id, productItem)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "Update failed",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusAccepted, res.Response{
		StatusCode: 200,
		Message:    "Updated",
		Data:       UpdateProductItem,
		Errors:     nil,
	})
}
func (cr *ProductHandler) DeleteProductItem(c *gin.Context) {
	paramsId := c.Param("id")
	id, err := strconv.Atoi(paramsId)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "can't find id",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	err = cr.productuseCase.DeleteProductItem(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "can't delete item",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, res.Response{
		StatusCode: 200,
		Message:    "item deleted",
		Data:       nil,
		Errors:     nil,
	})
}
func (cr *ProductHandler) DisaplyaAllProductItems(c *gin.Context) {
	paramsId := c.Param("id")
	id, err := strconv.Atoi(paramsId)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "can't find id",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}

	list, err := cr.productuseCase.DisaplyaAllProductItems(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "Failed to list",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusAccepted, res.Response{
		StatusCode: 200,
		Message:    "All-Product-items",
		Data:       list,
		Errors:     nil,
	})
}
