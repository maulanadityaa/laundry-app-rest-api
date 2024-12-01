package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/maulanadityaa/laundry-app-rest-api/helper"
	"github.com/maulanadityaa/laundry-app-rest-api/middleware"
	"github.com/maulanadityaa/laundry-app-rest-api/model/dto/request"
	"github.com/maulanadityaa/laundry-app-rest-api/model/dto/response"
	"github.com/maulanadityaa/laundry-app-rest-api/service"
	"github.com/maulanadityaa/laundry-app-rest-api/service/impl"
	"github.com/maulanadityaa/laundry-app-rest-api/validator"
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

// GetAllCustomer handles fetching all customers
// @Summary Get All Customers
// @Description Get all customers with pagination only for employees
// @Tags Customers
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param page query int false "Page"
// @Param rowsPerPage query int false "Rows Per Page"
// @Param name query string false "Name"
// @Success 200 {array} response.UserResponse
// @Router /api/v1/customers [get]
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

// GetCustomerByID handles fetching a customer by ID
// @Summary Get Customer By ID
// @Description Get a customer by ID only for employees and the customer itself
// @Tags Customers
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Customer ID"
// @Success 200 {object} response.UserResponse
// @Router /api/v1/customers/{id} [get]
func (CustomerController) GetCustomerByID(c *gin.Context) {
	id := c.Param("id")

	result, err := customerService.GetCustomerByID(id)
	if err != nil {
		response.NewResponseError(c, err.Error())
		return
	}

	response.NewResponseOK(c, result)
}

// UpdateCustomer handles updating a customer
// @Summary Update Customer
// @Description Update a customer only for employees and the customer itself
// @Tags Customers
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body request.UserUpdateRequest true "User Update Request Body"
// @Success 200 {object} response.UserResponse
// @Router /api/v1/customers [put]
func (CustomerController) UpdateCustomer(c *gin.Context) {
	var request request.UserUpdateRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		response.NewResponseBadRequest(c, err.Error())
		return
	}

	errors := validator.ValidateStruct(request)
	if errors != nil {
		response.NewResponseValidationError(c, errors)
		return
	}

	result, err := customerService.UpdateCustomer(request)
	if err != nil {
		response.NewResponseError(c, err.Error())
		return
	}

	response.NewResponseOK(c, result)
}

// GetCustomerByAccountID handles fetching a customer by account ID
// @Summary Get Customer By Account ID
// @Description Get a customer by account ID only for employees and the customer itself
// @Tags Customers
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param accountID path string true "Account ID"
// @Success 200 {object} response.UserResponse
// @Router /api/v1/customers/account/{accountID} [get]
func (CustomerController) GetCustomerByAccountID(c *gin.Context) {
	accountID := c.Param("accountID")

	result, err := customerService.GetCustomerByAccountID(accountID)
	if err != nil {
		response.NewResponseError(c, err.Error())
		return
	}

	response.NewResponseOK(c, result)
}
