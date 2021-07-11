package app

import (
	elasticsearch "bookstore_items-api/clients/elastic_search"

	"github.com/gorilla/mux"
)

const PORT = "8002"

var router = mux.NewRouter()

func StartApplication() {
	elasticsearch.Init()
	mapUrls()
	StartServer()
}
