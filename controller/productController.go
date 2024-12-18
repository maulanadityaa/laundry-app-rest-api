package controller

import (
	"maulanadityaa/laundry-app-rest-api/helper"
	"maulanadityaa/laundry-app-rest-api/middleware"
	"maulanadityaa/laundry-app-rest-api/model/dto/request"
	"maulanadityaa/laundry-app-rest-api/model/dto/response"
	"maulanadityaa/laundry-app-rest-api/service"
	"maulanadityaa/laundry-app-rest-api/service/impl"

	"github.com/gin-gonic/gin"
)

type ProductController struct{}

var productService service.ProductService = impl.NewProductService()

func NewProductController(g *gin.RouterGroup) {
	controller := new(ProductController)

	productGroup := g.Group("/products")
	{
		productGroup.POST("/", middleware.AuthWithRole([]string{"ROLE_EMPLOYEE"}), controller.AddProduct).Use(helper.ValidateJWT())
		productGroup.PUT("/", middleware.AuthWithRole([]string{"ROLE_EMPLOYEE"}), controller.UpdateProduct).Use(helper.ValidateJWT())
		productGroup.DELETE("/:id", middleware.AuthWithRole([]string{"ROLE_EMPLOYEE"}), controller.DeleteProduct).Use(helper.ValidateJWT())
		productGroup.GET("/", controller.GetAllProduct)
		productGroup.GET("/:id", controller.GetProductByID)
	}
}

func (ProductController) AddProduct(c *gin.Context) {
	var request request.ProductRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		response.NewResponseBadRequest(c, err.Error())
		return
	}

	result, err := productService.AddProduct(request)
	if err != nil {
		response.NewResponseError(c, err.Error())
		return
	}

	response.NewResponseCreated(c, result)
}

func (ProductController) UpdateProduct(c *gin.Context) {
	var request request.ProductUpdateRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		response.NewResponseBadRequest(c, err.Error())
		return
	}

	result, err := productService.UpdateProduct(request)
	if err != nil {
		response.NewResponseError(c, err.Error())
		return
	}

	response.NewResponseOK(c, result)
}

func (ProductController) DeleteProduct(c *gin.Context) {
	id := c.Param("id")

	err := productService.DeleteProduct(id)
	if err != nil {
		response.NewResponseError(c, err.Error())
		return
	}

	response.NewResponseOK(c, nil)
}

func (ProductController) GetAllProduct(c *gin.Context) {
	paging := c.DefaultQuery("paging", "1")
	rowsPerPage := c.DefaultQuery("rowsPerPage", "10")
	name := c.DefaultQuery("name", "")

	result, totalRows, TotalPage, err := productService.GetAllProduct(paging, rowsPerPage, name)
	if err != nil {
		response.NewResponseError(c, err.Error())
		return
	}

	response.NewResponseOKWithPaging(c, result, paging, rowsPerPage, totalRows, TotalPage)
}

func (ProductController) GetProductByID(c *gin.Context) {
	id := c.Param("id")

	result, err := productService.GetProductByID(id)
	if err != nil {
		response.NewResponseError(c, err.Error())
		return
	}

	response.NewResponseOK(c, result)
}
