package middleware

import (
	"maulanadityaa/laundry-app-rest-api/helper"
	"maulanadityaa/laundry-app-rest-api/model/dto/response"

	"github.com/gin-gonic/gin"
)

func AuthWithRole(role []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		jwtClaims := helper.GetJWTClaims(c)

		userRole := jwtClaims["role"].(string)

		for _, r := range role {
			if userRole == r {
				c.Next()
				return
			}
		}

		response.NewResponseForbidden(c, "You don't have permission to access this endpoint")
		c.Abort()
	}
}
