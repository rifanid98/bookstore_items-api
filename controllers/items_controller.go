package controllers

import (
	"bookstore_items-api/domain/items"
	"bookstore_items-api/services"
	"fmt"
	"net/http"

	"github.com/rifanid98/bookstore_oauth-go/oauth"
)

type IItemsController interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
}

type itemsController struct{}

var Items IItemsController = &itemsController{}

func (c *itemsController) Create(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthenticateRequest(r); err != nil {
		//TODO: return error to the user
		return
	}

	item := &items.Item{
		Seller: oauth.GetCallerId(r),
	}

	res, err := services.ItemService.Create(item)
	if err != nil {
		//TODO: return error json to the user
	}

	fmt.Println(res)
	//TODO: return 201
}

func (c *itemsController) Get(w http.ResponseWriter, r *http.Request) {

}
