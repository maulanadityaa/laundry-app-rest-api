package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/maulanadityaa/laundry-app-rest-api/middleware"
	"github.com/maulanadityaa/laundry-app-rest-api/model/dto/request"
	"github.com/maulanadityaa/laundry-app-rest-api/model/dto/response"
	"github.com/maulanadityaa/laundry-app-rest-api/service"
	"github.com/maulanadityaa/laundry-app-rest-api/service/impl"
)

type ProductController struct{}

var productService service.ProductService = impl.NewProductService()

func NewProductController(g *gin.RouterGroup) {
	controller := new(ProductController)

	productGroup := g.Group("/products")
	{
		productGroup.POST("/", middleware.AuthWithRole([]string{"ROLE_EMPLOYEE"}), controller.AddProduct)
		productGroup.PUT("/", middleware.AuthWithRole([]string{"ROLE_EMPLOYEE"}), controller.UpdateProduct)
		productGroup.DELETE("/:id", middleware.AuthWithRole([]string{"ROLE_EMPLOYEE"}), controller.DeleteProduct)
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
