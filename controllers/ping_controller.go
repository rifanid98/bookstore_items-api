package controllers

import (
	"bookstore_items-api/utils"
	"net/http"

	"github.com/rifanid98/bookstore_helper-go/response"
)

const pong = "PONG"

type IPingController interface {
	Ping(http.ResponseWriter, *http.Request)
}

type pingController struct{}

var Ping IPingController = &pingController{}

func (c *pingController) Ping(w http.ResponseWriter, r *http.Request) {
	utils.MyHttp.ToJsonRest(w, response.Success("Pong"), http.StatusOK)
}
