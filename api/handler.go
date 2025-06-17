package api

import (
	"net/http"

	"maulanadityaa/laundry-app-rest-api/app"
)

var router = app.InitApp()

func Handler(w http.ResponseWriter, r *http.Request) {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }

	router.ServeHTTP(w, r)
}
