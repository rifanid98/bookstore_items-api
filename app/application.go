package app

import (
	"github.com/gorilla/mux"
)

const PORT = "8002"

var router = mux.NewRouter()

func StartApplication() {
	mapUrls()
	StartServer()
}
