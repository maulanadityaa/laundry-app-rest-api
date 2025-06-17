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
		swaggerGroup.GET("/docs/*any", func(c *gin.Context) {
			path := c.Param("any")
			if path == "/" || path == "" {
				// Redirect /docs/ or /docs to index.html
				c.Redirect(301, "/api/v1/swagger/docs/index.html")
				return
			}

			// Otherwise, serve Swagger
			ginSwagger.WrapHandler(swaggerFiles.Handler)(c)
		})

	}
}
