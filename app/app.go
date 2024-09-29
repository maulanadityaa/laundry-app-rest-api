package app

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/maulanadityaa/laundry-app-rest-api/config"
)

func InitApp() {
	r := gin.Default()

	location, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		fmt.Println(err.Error())
	}

	time.Local = location

	config.LoadConfig()
	config.ConnectDB()

	addr := flag.String("port", ":"+os.Getenv("PORT"), "Address to listen and serve")
	if err := r.Run(*addr); err != nil {
		fmt.Println(err.Error())
	}
}
