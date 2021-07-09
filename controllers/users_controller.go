package controllers

import "net/http"

type IUsersController interface {
	Create(http.ResponseWriter, *http.Request)
}

type userController struct{}

var Users IUsersController = &userController{}

func (c *userController) Create(http.ResponseWriter, *http.Request) {

}
