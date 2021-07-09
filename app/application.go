package app

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

const PORT = "8002"

var router = mux.NewRouter()

func StartApplication() {
	mapUrls()

	srv := &http.Server{
		Handler: router,
		Addr:    fmt.Sprintf("localhost:%s", PORT),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	if err := srv.ListenAndServe(); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Server runinng on port: ", PORT)
	}
}
