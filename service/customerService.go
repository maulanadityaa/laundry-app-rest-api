package service

import (
	"github.com/maulanadityaa/laundry-app-rest-api/model/dto/request"
	"github.com/maulanadityaa/laundry-app-rest-api/model/dto/response"
)

type CustomerService interface {
	AddCustomer(req request.UserRequest) (response.UserResponse, error)
	UpdateCustomer(req request.UserUpdateRequest) (response.UserResponse, error)
	GetCustomerByID(customerID string) (response.UserResponse, error)
	GetCustomerByAccountID(accountID string) (response.UserResponse, error)
	GetAllCustomer(paging, rowsPerPage, name string) ([]response.UserResponse, string, string, error)
}
