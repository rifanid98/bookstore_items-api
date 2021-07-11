package controllers

import (
	"bookstore_items-api/domain/esqueries"
	"bookstore_items-api/domain/items"
	"bookstore_items-api/services"
	"bookstore_items-api/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	resp "github.com/rifanid98/bookstore_helper-go/response"
	"github.com/rifanid98/bookstore_oauth-go/oauth"
)

type IItemsController interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	Search(w http.ResponseWriter, r *http.Request)
	GetById(w http.ResponseWriter, r *http.Request)
}

type itemsController struct{}

var Items IItemsController = &itemsController{}

func (c *itemsController) Create(w http.ResponseWriter, r *http.Request) {
	if restErr := oauth.AuthenticateRequest(r); restErr != nil {
		utils.MyHttp.ToJsonRestErr(w, (*resp.RestErr)(restErr))
		return
	}

	sellerId := oauth.GetCallerId(r)
	if sellerId == 0 {
		utils.MyHttp.ToJsonRestErr(w, resp.Unauthorized(""))
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
		fmt.Println(err.Error())
		utils.MyHttp.ToJsonRestErr(w, resp.BadRequest("Invalid body"))
		return
	}

	item.Seller = sellerId

	res, restErr := services.Items.Create(&item)
	if err != nil {
		utils.MyHttp.ToJsonRestErr(w, restErr)
		return
	}

	utils.MyHttp.ToJsonRest(w, res, http.StatusCreated)
}

func (c *itemsController) Get(w http.ResponseWriter, r *http.Request) {

}

func (c *itemsController) Search(w http.ResponseWriter, r *http.Request) {
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		utils.MyHttp.ToJsonRestErr(w, resp.BadRequest("invalid json body"))
		return
	}
	defer r.Body.Close()

	query := &esqueries.EsQuery{}
	if err := json.Unmarshal(bytes, &query); err != nil {
		fmt.Println(err.Error())
		utils.MyHttp.ToJsonRestErr(w, resp.BadRequest("failed to unmarshal json"))
		return
	}

	items, restErr := services.Items.Search(query)
	if restErr != nil {
		utils.MyHttp.ToJsonRestErr(w, restErr)
		return
	}

	utils.MyHttp.ToJsonRestResp(w, resp.Success(items))
}

func (c *itemsController) GetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := strings.TrimSpace(vars["id"])

	item, err := services.Items.GetById(id)
	if err != nil {
		utils.MyHttp.ToJsonRestErr(w, err)
		return
	}

	utils.MyHttp.ToJsonRestResp(w, resp.Success(item))
}
