package services

import (
	"bookstore_items-api/domain/items"

	resp "github.com/rifanid98/bookstore_helper-go/response"
)

type IItemService interface {
	Create(*items.Item) (*items.Item, *resp.RestErr)
	Get(string) (*items.Item, *resp.RestErr)
}

type itemService struct {
}

var Items IItemService = &itemService{}

func (s *itemService) Create(item *items.Item) (*items.Item, *resp.RestErr) {
	item.Save()

	if err := item.Save(); err != nil {
		return nil, err
	}

	return item, nil
}

func (s *itemService) Get(a string) (*items.Item, *resp.RestErr) {
	return nil, resp.NotImplemented("")
}
