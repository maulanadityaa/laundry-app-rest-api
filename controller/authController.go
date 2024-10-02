package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/maulanadityaa/laundry-app-rest-api/model/dto/request"
	"github.com/maulanadityaa/laundry-app-rest-api/model/dto/response"
	"github.com/maulanadityaa/laundry-app-rest-api/service"
	"github.com/maulanadityaa/laundry-app-rest-api/service/impl"
	"github.com/maulanadityaa/laundry-app-rest-api/validator"
)

type AuthController struct{}

var authService service.AuthService = impl.NewAuthService()

func NewAuthController(g *gin.RouterGroup) {
	controller := new(AuthController)

	authGroup := g.Group("/auth")
	{
		authGroup.POST("/login", controller.Login)
		authGroup.POST("/register", controller.Register)
	}
}

func (AuthController) Login(c *gin.Context) {
	var request request.LoginRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		response.NewResponseBadRequest(c, err.Error())
		return
	}

	errors := validator.ValidateStruct(request)
	if errors != nil {
		response.NewResponseValidationError(c, errors)
		return
	}

	result, err := authService.Login(request)
	if err != nil {
		response.NewResponseError(c, err.Error())
		return
	}

	response.NewResponseOK(c, result)
}

func (AuthController) Register(c *gin.Context) {
	var request request.RegisterRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		response.NewResponseBadRequest(c, err.Error())
		return
	}

	errors := validator.ValidateStruct(request)
	if errors != nil {
		response.NewResponseValidationError(c, errors)
		return
	}

	result, err := authService.Register(request)
	if err != nil {
		response.NewResponseError(c, err.Error())
		return
	}

	response.NewResponseCreated(c, result)
}
