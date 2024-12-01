package main

import (
	"github.com/maulanadityaa/laundry-app-rest-api/app"
)

func main() {
	// @title Laundry App REST API
	// @version 1.0
	// @description This is a REST API application for laundry app
	// OpenAPI version specification
	// @OpenAPI 3.0.0

	// @contact.name maulanadityaa
	// @contact.url https://github.com/maulanadityaa
	// @contact.email maulanadityaaa@gmail.com

	// @schemes http https

	// Servers list
	// @servers [
	//   {
	//     "url": "http://localhost:8080",
	//     "description": "Local Development Server"
	//   },
	//   {
	//     "url": "https://api.yourproduction.com",
	//     "description": "Production Server"
	//   },
	//   {
	//     "url": "https://staging-api.yourcompany.com",
	//     "description": "Staging Server"
	//   }
	// ]

	// @securityDefinitions.apikey BearerAuth
	// @in header
	// @name Authorization
	// @description Type "Bearer" followed by a space and your token
	app.InitApp()
}
