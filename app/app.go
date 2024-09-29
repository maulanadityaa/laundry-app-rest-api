package app

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/maulanadityaa/laundry-app-rest-api/config"
	"github.com/maulanadityaa/laundry-app-rest-api/router"
)

func initDomainModule(r *gin.Engine) {
	apiGroup := r.Group("/api")
	v1Group := apiGroup.Group("/v1")

	router.InitRoutes(v1Group)
}

func InitApp() {
	r := gin.Default()

	location, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		fmt.Println(err.Error())
	}

	time.Local = location

	config.LoadConfig()
	config.ConnectDB()

	initDomainModule(r)

	addr := flag.String("port", ":"+os.Getenv("PORT"), "Address to listen and serve")
	if err := r.Run(*addr); err != nil {
		fmt.Println(err.Error())
	}
}
