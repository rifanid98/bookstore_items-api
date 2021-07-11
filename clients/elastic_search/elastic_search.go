package elasticsearch

import (
	"context"
	"fmt"
	"time"

	"github.com/olivere/elastic/v7"
	"github.com/rifanid98/bookstore_utils-go/logger"
)

type IClient interface {
	setClient(*elastic.Client)
	Index(string, interface{}) (*elastic.IndexResponse, error)
	Get(string, string) (*elastic.GetResult, error)
}

type esClient struct {
	client *elastic.Client
}

var Client = &esClient{}

func Init() {
	log := logger.GetLogger()

	client, err := elastic.NewClient(
		elastic.SetURL("http://localhost:9200"),
		elastic.SetHealthcheckInterval(10*time.Second),
		elastic.SetErrorLog(log),
		elastic.SetInfoLog(log),
	)

	if err != nil {
		panic(err)
	}

	Client.setClient(client)
}

func (c *esClient) setClient(client *elastic.Client) {
	c.client = client
}

func (c *esClient) Index(index string, doc interface{}) (*elastic.IndexResponse, error) {
	ctx := context.Background()
	res, err := c.client.Index().
		Index(index).
		Type("_doc").
		BodyJson(doc).
		Do(ctx)

	if err != nil {
		logger.Error(
			fmt.Sprintf("error when trying to index document in %s", index),
			err,
		)
		return nil, err
	}

	return res, nil
}

func (c *esClient) Get(index string) (*elastic.GetResult, error) {
	res, err := c.client.Get().
		Index(index).
		Type("_doc").
		Do(context.Background())

	if err != nil {
		logger.Error(fmt.Sprintf("error when trying to get data"), err)
		return nil, err
	}

	if !res.Found {
		return nil, nil
	}

	return res, nil
}

func (c *esClient) GetById(index string, id string) (*elastic.GetResult, error) {
	res, err := c.client.Get().
		Index(index).
		Type("_doc").
		Id(id).
		Do(context.Background())

	if err != nil {
		logger.Error(fmt.Sprintf("error when trying to get id %s", id), err)
		return nil, err
	}

	if !res.Found {
		return nil, nil
	}

	return res, nil
}
