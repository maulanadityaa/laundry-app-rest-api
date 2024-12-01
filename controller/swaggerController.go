package controller

import (
	"github.com/gin-gonic/gin"
	_ "github.com/maulanadityaa/laundry-app-rest-api/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type SwaggerController struct{}

func NewSwaggerController(route *gin.RouterGroup) {
	swaggerGroup := route.Group("/swagger")
	{
		swaggerGroup.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("http://localhost:8080/api/v1/swagger/docs/doc.json"), ginSwagger.DefaultModelsExpandDepth(-1)))
	}
}
