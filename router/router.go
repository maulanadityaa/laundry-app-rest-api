package router

import (
	"github.com/gin-gonic/gin"
	"github.com/maulanadityaa/laundry-app-rest-api/controller"
)

func InitRoutes(route *gin.RouterGroup) {
	controller.NewAuthController(route)
	controller.NewCustomerController(route)
	controller.NewEmployeeController(route)
}
