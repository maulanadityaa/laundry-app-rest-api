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

// Login handles user login
// @Summary User Login
// @Description Authenticate a user and return a JWT token
// @Tags Authentication
// @Accept json
// @Produce json
// @Param request body request.LoginRequest true "Login Request Body"
// @Success 200 {object} response.LoginResponse
// @Router /api/v1/auth/login [post]
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

// Register handles user registration
// @Summary User Registration
// @Description Register a new customer or employee
// @Tags Authentication
// @Accept json
// @Produce json
// @Param request body request.RegisterRequest true "Registration Request Body"
// @Success 201 {object} response.RegisterResponse
// @Router /api/v1/auth/register [post]
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
