package controller

import (
	_ "maulanadityaa/laundry-app-rest-api/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type SwaggerController struct{}

func NewSwaggerController(route *gin.RouterGroup) {
	swaggerGroup := route.Group("/swagger")
	{
		swaggerGroup.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.DefaultModelsExpandDepth(-1)))
	}
}
