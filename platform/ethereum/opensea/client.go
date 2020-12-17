package opensea

import (
	"net/url"
	"strconv"

	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

type Client struct {
	blockatlas.Request
}

func InitClient(api string, apiKey string) *Client {
	c := Client{blockatlas.InitClient(api)}
	c.Headers["X-API-KEY"] = apiKey
	return &c
}

func (c Client) getCollections(owner string) (page []Collection, err error) {
	query := url.Values{
		"asset_owner": {owner},
		"limit":       {"1000"},
	}
	err = c.Get(&page, "api/v1/collections", query)
	return
}

func (c Client) getCollectibles(owner string, collectionID string) ([]Collectible, error) {
	query := url.Values{
		"owner":      {owner},
		"collection": {collectionID},
		"limit":      {strconv.Itoa(300)},
	}

	var page CollectiblePage
	err := c.Get(&page, "api/v1/assets", query)
	return page.Collectibles, err
}
