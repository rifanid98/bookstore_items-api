package items

import (
	elasticsearch "bookstore_items-api/clients/elastic_search"

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
