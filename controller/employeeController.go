package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/maulanadityaa/laundry-app-rest-api/model/dto/request"
	"github.com/maulanadityaa/laundry-app-rest-api/model/dto/response"
	"github.com/maulanadityaa/laundry-app-rest-api/service"
	"github.com/maulanadityaa/laundry-app-rest-api/service/impl"
)

type EmployeeController struct{}

var employeeService service.EmployeeService = impl.NewEmployeeService()

func NewEmployeeController(g *gin.RouterGroup) {
	controller := new(EmployeeController)

	employeeGroup := g.Group("/employees")
	{
		employeeGroup.GET("/", controller.GetAllEmployee)
		employeeGroup.GET("/:id", controller.GetEmployeeByID)
		employeeGroup.PUT("/", controller.UpdateEmployee)
		employeeGroup.GET("/account/:accountID", controller.GetEmployeeByAccountID)
	}
}

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

func (EmployeeController) GetEmployeeByID(c *gin.Context) {
	id := c.Param("id")

	result, err := employeeService.GetEmployeeByID(id)
	if err != nil {
		response.NewResponseError(c, err.Error())
		return
	}

	response.NewResponseOK(c, result)
}

func (EmployeeController) UpdateEmployee(c *gin.Context) {
	var request request.UserUpdateRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		response.NewResponseBadRequest(c, err.Error())
		return
	}

	result, err := employeeService.UpdateEmployee(request)
	if err != nil {
		response.NewResponseError(c, err.Error())
		return
	}

	response.NewResponseOK(c, result)
}

func (EmployeeController) GetEmployeeByAccountID(c *gin.Context) {
	accountID := c.Param("accountID")

	result, err := employeeService.GetEmployeeByAccountID(accountID)
	if err != nil {
		response.NewResponseError(c, err.Error())
		return
	}

	response.NewResponseOK(c, result)
}
