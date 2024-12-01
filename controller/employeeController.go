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

type EmployeeController struct{}

var employeeService service.EmployeeService = impl.NewEmployeeService()

func NewEmployeeController(g *gin.RouterGroup) {
	controller := new(EmployeeController)

	employeeGroup := g.Group("/employees", helper.ValidateJWT())
	employeeGroup.Use(middleware.AuthWithRole([]string{"ROLE_EMPLOYEE"}))
	{
		employeeGroup.GET("/", controller.GetAllEmployee)
		employeeGroup.GET("/:id", controller.GetEmployeeByID)
		employeeGroup.PUT("/", controller.UpdateEmployee)
		employeeGroup.GET("/account/:accountID", controller.GetEmployeeByAccountID)
	}
}

// GetAllEmployee handles fetching all employees
// @Summary Get All Employees
// @Description Get all employees with pagination only for employees
// @Tags Employees
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param page query int false "Page"
// @Param rowsPerPage query int false "Rows Per Page"
// @Param name query string false "Name"
// @Success 200 {array} response.UserResponse
// @Router /api/v1/employees [get]
func (EmployeeController) GetAllEmployee(c *gin.Context) {
	paging := c.DefaultQuery("page", "1")
	rowsPerPage := c.DefaultQuery("rowsPerPage", "10")
	name := c.DefaultQuery("name", "")

	result, totalRows, totalPage, err := employeeService.GetAllEmployee(paging, rowsPerPage, name)
	if err != nil {
		response.NewResponseError(c, err.Error())
		return
	}

	response.NewResponseOKWithPaging(c, result, paging, rowsPerPage, totalRows, totalPage)
}

// GetEmployeeByID handles fetching an employee by ID
// @Summary Get Employee By ID
// @Description Get an employee by ID only for employees
// @Tags Employees
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Employee ID"
// @Success 200 {object} response.UserResponse
// @Router /api/v1/employees/{id} [get]
func (EmployeeController) GetEmployeeByID(c *gin.Context) {
	id := c.Param("id")

	result, err := employeeService.GetEmployeeByID(id)
	if err != nil {
		response.NewResponseError(c, err.Error())
		return
	}

	response.NewResponseOK(c, result)
}

// UpdateEmployee handles updating an employee
// @Summary Update Employee
// @Description Update an employee only for employees
// @Tags Employees
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body request.UserUpdateRequest true "User Update Request Body"
// @Success 200 {object} response.UserResponse
// @Router /api/v1/employees [put]
func (EmployeeController) UpdateEmployee(c *gin.Context) {
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

	result, err := employeeService.UpdateEmployee(request)
	if err != nil {
		response.NewResponseError(c, err.Error())
		return
	}

	response.NewResponseOK(c, result)
}

// GetEmployeeByAccountID handles fetching an employee by account ID
// @Summary Get Employee By Account ID
// @Description Get an employee by account ID only for employees
// @Tags Employees
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param accountID path string true "Account ID"
// @Success 200 {object} response.UserResponse
// @Router /api/v1/employees/account/{accountID} [get]
func (EmployeeController) GetEmployeeByAccountID(c *gin.Context) {
	accountID := c.Param("accountID")

	result, err := employeeService.GetEmployeeByAccountID(accountID)
	if err != nil {
		response.NewResponseError(c, err.Error())
		return
	}

	response.NewResponseOK(c, result)
}
