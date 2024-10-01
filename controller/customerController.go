package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/maulanadityaa/laundry-app-rest-api/helper"
	"github.com/maulanadityaa/laundry-app-rest-api/middleware"
	"github.com/maulanadityaa/laundry-app-rest-api/model/dto/request"
	"github.com/maulanadityaa/laundry-app-rest-api/model/dto/response"
	"github.com/maulanadityaa/laundry-app-rest-api/service"
	"github.com/maulanadityaa/laundry-app-rest-api/service/impl"
)

type CustomerController struct{}

var customerService service.CustomerService = impl.NewCustomerService()

func NewCustomerController(g *gin.RouterGroup) {
	controller := new(CustomerController)

	customerGroup := g.Group("/customers", helper.ValidateJWT())
	{
		customerGroup.GET("/", middleware.AuthWithRole([]string{"ROLE_EMPLOYEE"}), controller.GetAllCustomer)
		customerGroup.GET("/:id", middleware.AuthWithRole([]string{"ROLE_EMPLOYEE", "ROLE_CUSTOMER"}), controller.GetCustomerByID)
		customerGroup.PUT("/", middleware.AuthWithRole([]string{"ROLE_EMPLOYEE", "ROLE_CUSTOMER"}), controller.UpdateCustomer)
		customerGroup.GET("/account/:accountID", middleware.AuthWithRole([]string{"ROLE_EMPLOYEE", "ROLE_CUSTOMER"}), controller.GetCustomerByAccountID)
	}
}

func (CustomerController) GetAllCustomer(c *gin.Context) {
	paging := c.DefaultQuery("page", "1")
	rowsPerPage := c.DefaultQuery("rowsPerPage", "10")
	name := c.DefaultQuery("name", "")

	result, totalRows, totalPage, err := customerService.GetAllCustomer(paging, rowsPerPage, name)
	if err != nil {
		response.NewResponseError(c, err.Error())
		return
	}

	response.NewResponseOKWithPaging(c, result, paging, rowsPerPage, totalRows, totalPage)
}

func (CustomerController) GetCustomerByID(c *gin.Context) {
	id := c.Param("id")

	result, err := customerService.GetCustomerByID(id)
	if err != nil {
		response.NewResponseError(c, err.Error())
		return
	}

	response.NewResponseOK(c, result)
}

func (CustomerController) UpdateCustomer(c *gin.Context) {
	var request request.UserUpdateRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		response.NewResponseBadRequest(c, err.Error())
		return
	}

	result, err := customerService.UpdateCustomer(request)
	if err != nil {
		response.NewResponseError(c, err.Error())
		return
	}

	response.NewResponseOK(c, result)
}

func (CustomerController) GetCustomerByAccountID(c *gin.Context) {
	accountID := c.Param("accountID")

	result, err := customerService.GetCustomerByAccountID(accountID)
	if err != nil {
		response.NewResponseError(c, err.Error())
		return
	}

	response.NewResponseOK(c, result)
}
