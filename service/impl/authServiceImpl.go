package impl

import (
	"errors"
	"fmt"
	"time"

	"maulanadityaa/laundry-app-rest-api/helper"
	"maulanadityaa/laundry-app-rest-api/model/dto/request"
	"maulanadityaa/laundry-app-rest-api/model/dto/response"
	"maulanadityaa/laundry-app-rest-api/model/entity"
	"maulanadityaa/laundry-app-rest-api/repository"
	"maulanadityaa/laundry-app-rest-api/repository/impl"
	"maulanadityaa/laundry-app-rest-api/service"

	"github.com/google/uuid"
)

type AuthService struct{}

var accountRepository repository.AccountRepository = impl.NewAccountRepository()
var roleRepository repository.RoleRepository = impl.NewRoleRepository()
var customerService service.CustomerService = NewCustomerService()
var employeeService service.EmployeeService = NewEmployeeService()

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (AuthService) Register(req request.RegisterRequest) (response.RegisterResponse, error) {
	role, _ := roleRepository.GetRoleByName(req.Role)
	hashedPassword, _ := helper.HashPassword(req.Password)
	fmt.Println(role)

	newAccount := entity.Account{
		ID:        uuid.NewString(),
		Email:     req.Email,
		Password:  hashedPassword,
		RoleID:    role.ID,
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	account, err := accountRepository.AddAccount(newAccount)
	if err != nil {
		return response.RegisterResponse{}, err
	}

	newUserRequest := request.UserRequest{
		Name:        req.Name,
		Address:     req.Address,
		PhoneNumber: helper.FormatPhoneNumber(req.PhoneNumber),
		AccountID:   account.ID,
	}

	var userResponse response.UserResponse

	if role.Name == "ROLE_CUSTOMER" {
		userResponse, _ = customerService.AddCustomer(newUserRequest)
	} else if role.Name == "ROLE_EMPLOYEE" {
		userResponse, _ = employeeService.AddEmployee(newUserRequest)
	} else {
		return response.RegisterResponse{}, errors.New("role not found")
	}

	return response.RegisterResponse{
		ID:    account.ID,
		Email: account.Email,
		Role: response.RoleResponse{
			ID:   role.ID,
			Name: role.Name,
		},
		UserResponse: response.UserResponse{
			ID:          userResponse.ID,
			Name:        userResponse.Name,
			Address:     userResponse.Address,
			PhoneNumber: userResponse.PhoneNumber,
			CreatedAt:   userResponse.CreatedAt,
			UpdatedAt:   userResponse.UpdatedAt,
		},
	}, nil
}

func (AuthService) Login(req request.LoginRequest) (response.LoginResponse, error) {
	account, err := accountRepository.GetAccountByEmail(req.Email)
	if err != nil {
		return response.LoginResponse{}, err
	}

	err = helper.ComparePassword(account.Password, req.Password)
	if err != nil {
		return response.LoginResponse{}, err
	}

	role, err := roleRepository.GetRoleByID(account.RoleID)
	if err != nil {
		return response.LoginResponse{}, err
	}

	token, err := helper.GenerateJWT(account.ID, role.Name, account.Email)
	if err != nil {
		return response.LoginResponse{}, err
	}

	return response.LoginResponse{
		Token: token,
	}, nil
}
