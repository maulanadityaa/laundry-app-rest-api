package controller

import (
	_ "maulanadityaa/laundry-app-rest-api/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type SwaggerController struct{}

func NewSwaggerController(route *gin.Engine) {
	swaggerGroup := route.Group("/swagger")
	{
		// Redirect /swagger/docs → /swagger/docs/index.html
		swaggerGroup.GET("/docs", func(c *gin.Context) {
			c.Redirect(302, "/api/v1/swagger/docs/index.html")
		})

		// Serve Swagger UI
		swaggerGroup.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
}
