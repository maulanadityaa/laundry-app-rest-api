package impl

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/maulanadityaa/laundry-app-rest-api/helper"
	"github.com/maulanadityaa/laundry-app-rest-api/model/dto/request"
	"github.com/maulanadityaa/laundry-app-rest-api/model/dto/response"
	"github.com/maulanadityaa/laundry-app-rest-api/model/entity"
	"github.com/maulanadityaa/laundry-app-rest-api/repository"
	"github.com/maulanadityaa/laundry-app-rest-api/repository/impl"
	"gorm.io/gorm"
)

type CustomerService struct{}

var customerRepository repository.CustomerRepository = impl.NewCustomerRepository()

func NewCustomerService() *CustomerService {
	return &CustomerService{}
}

func (CustomerService) AddCustomer(req request.UserRequest) (response.UserResponse, error) {
	newCustomer := entity.Customer{
		ID:          uuid.NewString(),
		Name:        req.Name,
		Address:     req.Address,
		PhoneNumber: req.PhoneNumber,
		AccountID:   req.AccountID,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	customer, err := customerRepository.AddCustomer(newCustomer)
	if err != nil {
		return response.UserResponse{}, err
	}

	return response.UserResponse{
		ID:          customer.ID,
		Name:        customer.Name,
		Address:     customer.Address,
		PhoneNumber: customer.PhoneNumber,
		CreatedAt:   customer.CreatedAt.String(),
		UpdatedAt:   customer.UpdatedAt.String(),
	}, nil
}

func (CustomerService) UpdateCustomer(req request.UserUpdateRequest) (response.UserResponse, error) {
	customer, err := customerRepository.GetCustomerByID(req.ID)
	if err != nil {
		return response.UserResponse{}, err
	}

	customer.Name = req.Name
	customer.Address = req.Address
	customer.PhoneNumber = req.PhoneNumber
	customer.UpdatedAt = time.Now()

	customer, err = customerRepository.UpdateCustomer(customer)
	if err != nil {
		return response.UserResponse{}, err
	}

	return response.UserResponse{
		ID:          customer.ID,
		Name:        customer.Name,
		Address:     customer.Address,
		PhoneNumber: customer.PhoneNumber,
		CreatedAt:   customer.CreatedAt.String(),
		UpdatedAt:   customer.UpdatedAt.String(),
	}, nil
}

func (CustomerService) GetCustomerByID(customerID string) (response.UserResponse, error) {
	customer, err := customerRepository.GetCustomerByID(customerID)
	if err != nil {
		return response.UserResponse{}, err
	}

	return response.UserResponse{
		ID:          customer.ID,
		Name:        customer.Name,
		Address:     customer.Address,
		PhoneNumber: customer.PhoneNumber,
		CreatedAt:   customer.CreatedAt.String(),
		UpdatedAt:   customer.UpdatedAt.String(),
	}, nil
}

func (CustomerService) GetCustomerByAccountID(accountID string) (response.UserResponse, error) {
	customer, err := customerRepository.GetCustomerByAccountID(accountID)
	if err != nil {
		return response.UserResponse{}, err
	}

	return response.UserResponse{
		ID:          customer.ID,
		Name:        customer.Name,
		Address:     customer.Address,
		PhoneNumber: customer.PhoneNumber,
		CreatedAt:   customer.CreatedAt.String(),
		UpdatedAt:   customer.UpdatedAt.String(),
	}, nil
}

func (CustomerService) GetAllCustomer(paging, rowsPerPage, name string) ([]response.UserResponse, string, string, error) {
	pagingInt, err := strconv.Atoi(paging)
	if err != nil {
		fmt.Println(err)
		return nil, "0", "0", errors.New("invalid query parameter")
	}

	rowsPerPageInt, err := strconv.Atoi(rowsPerPage)
	if err != nil {
		fmt.Println(err)
		return nil, "0", "0", errors.New("invalid query parameter")
	}

	var spec []func(db *gorm.DB) *gorm.DB
	spec = append(spec, helper.Paginate(pagingInt, rowsPerPageInt))

	customers, totalRows, err := customerRepository.GetAllCustomer(spec, name)
	if err != nil {
		fmt.Println(err)
		return nil, "0", "0", err
	}

	customerResponses := make([]response.UserResponse, 0)
	for _, customer := range customers {
		customerResponses = append(customerResponses, response.UserResponse{
			ID:          customer.ID,
			Name:        customer.Name,
			PhoneNumber: customer.PhoneNumber,
			Address:     customer.Address,
			CreatedAt:   customer.CreatedAt.String(),
			UpdatedAt:   customer.UpdatedAt.String(),
		})
	}

	totalPage := helper.GetTotalPage(totalRows, rowsPerPageInt)

	return customerResponses, totalRows, strconv.Itoa(totalPage), nil
}
