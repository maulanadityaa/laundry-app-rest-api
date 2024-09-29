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

type EmployeeService struct{}

var employeeRepository repository.EmployeeRepository = impl.NewEmployeeRepository()

func NewEmployeeService() *EmployeeService {
	return &EmployeeService{}
}

func (EmployeeService) AddEmployee(req request.UserRequest) (response.UserResponse, error) {
	newEmployee := entity.Employee{
		ID:          uuid.NewString(),
		Name:        req.Name,
		Address:     req.Address,
		PhoneNumber: req.PhoneNumber,
		AccountID:   req.AccountID,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	employee, err := employeeRepository.AddEmployee(newEmployee)
	if err != nil {
		return response.UserResponse{}, err
	}

	return response.UserResponse{
		ID:          employee.ID,
		Name:        employee.Name,
		Address:     employee.Address,
		PhoneNumber: employee.PhoneNumber,
		CreatedAt:   employee.CreatedAt.String(),
		UpdatedAt:   employee.UpdatedAt.String(),
	}, nil
}

func (EmployeeService) UpdateEmployee(req request.UserUpdateRequest) (response.UserResponse, error) {
	employee, err := employeeRepository.GetEmployeeByID(req.ID)
	if err != nil {
		return response.UserResponse{}, err
	}

	employee.Name = req.Name
	employee.Address = req.Address
	employee.PhoneNumber = req.PhoneNumber
	employee.UpdatedAt = time.Now()

	employee, err = employeeRepository.UpdateEmployee(employee)
	if err != nil {
		return response.UserResponse{}, err
	}

	return response.UserResponse{
		ID:          employee.ID,
		Name:        employee.Name,
		Address:     employee.Address,
		PhoneNumber: employee.PhoneNumber,
		CreatedAt:   employee.CreatedAt.String(),
		UpdatedAt:   employee.UpdatedAt.String(),
	}, nil
}

func (EmployeeService) GetEmployeeByID(employeeID string) (response.UserResponse, error) {
	employee, err := employeeRepository.GetEmployeeByID(employeeID)
	if err != nil {
		return response.UserResponse{}, err
	}

	return response.UserResponse{
		ID:          employee.ID,
		Name:        employee.Name,
		Address:     employee.Address,
		PhoneNumber: employee.PhoneNumber,
		CreatedAt:   employee.CreatedAt.String(),
		UpdatedAt:   employee.UpdatedAt.String(),
	}, nil
}

func (EmployeeService) GetEmployeeByAccountID(accountID string) (response.UserResponse, error) {
	employee, err := employeeRepository.GetEmployeeByAccountID(accountID)
	if err != nil {
		return response.UserResponse{}, err
	}

	return response.UserResponse{
		ID:          employee.ID,
		Name:        employee.Name,
		Address:     employee.Address,
		PhoneNumber: employee.PhoneNumber,
		CreatedAt:   employee.CreatedAt.String(),
		UpdatedAt:   employee.UpdatedAt.String(),
	}, nil
}

func (EmployeeService) GetAllEmployee(paging, rowsPerPage, name string) ([]response.UserResponse, string, string, error) {
	pagingInt, err := strconv.Atoi(paging)
	if err != nil {
		return []response.UserResponse{}, "", "", errors.New("invalid query parameter")
	}

	rowsPerPageInt, err := strconv.Atoi(rowsPerPage)
	if err != nil {
		return []response.UserResponse{}, "", "", errors.New("invalid query parameter")
	}

	var spec []func(db *gorm.DB) *gorm.DB
	spec = append(spec, helper.Paginate(pagingInt, rowsPerPageInt))

	employees, totalRows, err := employeeRepository.GetAllEmployee(spec, name)
	if err != nil {
		fmt.Println(err)
		return []response.UserResponse{}, "", "", err
	}

	employeesResponse := []response.UserResponse{}
	for _, employee := range employees {
		employeesResponse = append(employeesResponse, response.UserResponse{
			ID:          employee.ID,
			Name:        employee.Name,
			PhoneNumber: employee.PhoneNumber,
			Address:     employee.Address,
			CreatedAt:   employee.CreatedAt.String(),
			UpdatedAt:   employee.UpdatedAt.String(),
		})
	}

	totalPage := helper.GetTotalPage(totalRows, rowsPerPageInt)

	return employeesResponse, totalRows, strconv.Itoa(totalPage), nil
}
