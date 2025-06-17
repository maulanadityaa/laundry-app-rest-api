// @title Laundry App REST API
// @version 1.0
// @description This is a REST API application for laundry app

// @contact.name maulanadityaa
// @contact.url https://github.com/maulanadityaa
// @contact.email maulanadityaaa@gmail.com

// @schemes http https

// @server http://localhost:8080 Local Development Server
// @server https://api.yourproduction.com Production Server
// @server https://staging-api.yourcompany.com Staging Server

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and your token

package main

import (
	"fmt"
	"log"
	"maulanadityaa/laundry-app-rest-api/app"
	"os"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	app := app.InitApp()

	fmt.Println("🟢 Starting server on", port)
	log.Printf("🟢 Server is running on port %s", port)

	if err := app.Run(": " + port); err != nil {
		log.Fatalf("🔥 Server failed to start: %v", err)
	}
}
