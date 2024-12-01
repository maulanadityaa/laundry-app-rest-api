package api

import (
	"net/http"

	"github.com/maulanadityaa/laundry-app-rest-api/app"
)

var router = app.InitApp()

func Handler(w http.ResponseWriter, r *http.Request) {
	router.ServeHTTP(w, r)
}
