package app

import (
	"bookstore_items-api/controllers"
	"net/http"
)

func mapUrls() {
	router.HandleFunc("/ping", controllers.Ping.Ping)
	router.HandleFunc("/items", controllers.Items.Create).Methods(http.MethodPost)
	router.HandleFunc("/items", controllers.Items.Create).Methods(http.MethodGet)
}