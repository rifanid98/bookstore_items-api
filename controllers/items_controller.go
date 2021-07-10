package controllers

import (
	"bookstore_items-api/domain/items"
	"bookstore_items-api/services"
	"bookstore_items-api/utils"
	"encoding/json"
	"io/ioutil"
	"net/http"

	resp "github.com/rifanid98/bookstore_helper-go/response"
	"github.com/rifanid98/bookstore_oauth-go/oauth"
)

type IItemsController interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
}

type itemsController struct{}

var Items IItemsController = &itemsController{}

func (c *itemsController) Create(w http.ResponseWriter, r *http.Request) {
	if restErr := oauth.AuthenticateRequest(r); restErr != nil {
		utils.MyHttp.ToJsonRestErr(w, (*resp.RestErr)(restErr))
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.MyHttp.ToJsonRestErr(w, resp.InternalServerError("Failed to read body"))
		return
	}
	defer r.Body.Close()

	var item items.Item
	if err := json.Unmarshal(body, &item); err != nil {
		utils.MyHttp.ToJsonRestErr(w, resp.BadRequest("Invalid body"))
		return
	}

	item.Seller = oauth.GetCallerId(r)

	res, restErr := services.Items.Create(&item)
	if err != nil {
		utils.MyHttp.ToJsonRestErr(w, restErr)
		return
	}

	utils.MyHttp.ToJsonRest(w, res, http.StatusCreated)
}

func (c *itemsController) Get(w http.ResponseWriter, r *http.Request) {

}
