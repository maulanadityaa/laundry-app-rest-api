package service

import (
	"github.com/maulanadityaa/laundry-app-rest-api/model/dto/request"
	"github.com/maulanadityaa/laundry-app-rest-api/model/dto/response"
)

type EmployeeService interface {
	AddEmployee(req request.UserRequest) (response.UserResponse, error)
	UpdateEmployee(req request.UserUpdateRequest) (response.UserResponse, error)
	GetEmployeeByID(employeeID string) (response.UserResponse, error)
	GetEmployeeByAccountID(accountID string) (response.UserResponse, error)
	GetAllEmployee(paging, rowsPerPage, name string) ([]response.UserResponse, string, string, error)
}
