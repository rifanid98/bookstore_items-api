package app

import (
	"fmt"
	"net/http"
	"time"
)

func StartServer() {
	srv := &http.Server{
		Handler: router,
		Addr:    fmt.Sprintf("localhost:%s", PORT),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 500 * time.Second,
		ReadTimeout:  2 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	if err := srv.ListenAndServe(); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Server runinng on port: ", PORT)
	}
}
