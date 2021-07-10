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

func (s *itemService) Create(*items.Item) (*items.Item, *resp.RestErr) {
	return nil, resp.NotImplemented("")
}

func (s *itemService) Get(a string) (*items.Item, *resp.RestErr) {
	return nil, resp.NotImplemented("")
}
