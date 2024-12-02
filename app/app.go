package app

import (
	"flag"
	"fmt"
	"log"
	"os"
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
	// gin.SetMode(gin.ReleaseMode)

	location, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		fmt.Println(err.Error())
	}
	log.Printf("Location: %s", location.String())

	time.Local = location

	go func() {
		config.LoadConfig()
		config.ConnectDB()
		validator.InitValidator()
	}()

	initDomainModule(r)

	addr := flag.String("port", ":"+os.Getenv("PORT"), "Address to listen and serve")
	if err := r.Run(*addr); err != nil {
		fmt.Println(err.Error())
	}
	log.Printf("Server is running on %s", os.Getenv("PORT"))

	return r
}
