package app

import (
	"log"
	"time"

	"maulanadityaa/laundry-app-rest-api/config"
	"maulanadityaa/laundry-app-rest-api/router"
	"maulanadityaa/laundry-app-rest-api/validator"

	"github.com/gin-gonic/gin"
)

func initDomainModule(r *gin.Engine) {
	apiGroup := r.Group("/api")
	v1Group := apiGroup.Group("/v1")

	router.InitRoutes(v1Group)
}

func InitApp() *gin.Engine {
	r := gin.Default()

	// Set time zone
	location, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		log.Fatal("Failed to load location:", err)
	}
	time.Local = location
	log.Printf("Timezone set to: %s", location.String())

	// Load essential components (not in goroutine)
	config.LoadConfig()
	config.ConnectDB()
	validator.InitValidator()

	// Register routes/modules
	initDomainModule(r)

	return r
}
