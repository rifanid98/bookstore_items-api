package controllers

import "net/http"

const pong = "PONG"

type IPingController interface {
	Ping(http.ResponseWriter, *http.Request)
}

type pingController struct{}

var Ping IPingController = &pingController{}

func (c *pingController) Ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(pong))
}
