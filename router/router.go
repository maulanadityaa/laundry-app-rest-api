package router

import (
	"maulanadityaa/laundry-app-rest-api/controller"

	"github.com/gin-gonic/gin"
)

func InitRoutes(route *gin.RouterGroup) {
	controller.NewAuthController(route)
	controller.NewCustomerController(route)
	controller.NewEmployeeController(route)
	controller.NewProductController(route)
	controller.NewTransactionController(route)
	controller.NewSwaggerController(route)
}
