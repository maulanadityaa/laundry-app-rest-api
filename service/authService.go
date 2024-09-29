package service

import (
	"github.com/maulanadityaa/laundry-app-rest-api/model/dto/request"
	"github.com/maulanadityaa/laundry-app-rest-api/model/dto/response"
)

type AuthService interface {
	Login(req request.LoginRequest) (response.LoginResponse, error)
	Register(req request.RegisterRequest) (response.RegisterResponse, error)
}
