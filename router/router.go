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

}

func InitSwaggerRoutes(route *gin.Engine) {
	controller.NewSwaggerController(route)
}
