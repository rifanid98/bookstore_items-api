package items

import (
	elasticsearch "bookstore_items-api/clients/elastic_search"
	"encoding/json"
	"fmt"

	"github.com/olivere/elastic/v7"
	resp "github.com/rifanid98/bookstore_helper-go/response"
)

const (
	indexItems = "items"
)

func (i *Item) Save() *resp.RestErr {
	res, err := elasticsearch.Client.Index(indexItems, i)
	if err != nil {
		return resp.InternalServerError("error when trying to save item")
	}

	i.Id = res.Id
	return nil
}

func (i *Item) Get() *resp.RestErr {
	res, err := elasticsearch.Client.Get(indexItems)
	if err != nil {
		return resp.InternalServerError(fmt.Sprintf("error when trying to get id %s", i.Id))
	}
	if !res.Found {
		return resp.NotFound(fmt.Sprintf("no item found with id %s", i.Id))
	}
	return nil
}

func (i *Item) GetById() *resp.RestErr {
	res, err := elasticsearch.Client.GetById(indexItems, i.Id)
	fmt.Println()
	if err != nil {
		if elastic.IsNotFound(err) {
			return resp.NotFound(fmt.Sprintf("no data found with id %s", i.Id))
		}
		return resp.InternalServerError(fmt.Sprintf("error when trying to get id %s", i.Id))
	}

	bytes, err := res.Source.MarshalJSON()
	if err != nil {
		fmt.Println(err.Error())
		return resp.InternalServerError("error when trying to marshal json")
	}
	if err := json.Unmarshal(bytes, i); err != nil {
		return resp.InternalServerError("error when trying to unmarshal json")
	}

	return nil
}
