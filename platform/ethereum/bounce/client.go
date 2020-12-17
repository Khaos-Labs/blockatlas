package bounce

import (
	"net/url"
	"strconv"

	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

type Client struct {
	blockatlas.Request
}

func (c Client) getCollections(owner string, chainId int) ([]Collection, error) {
	query := url.Values{
		"address":  {owner},
		"chain_id": {strconv.Itoa(chainId)},
	}
	var resp CollectionResponse
	err := c.Get(&resp, "/nft", query)
	if err != nil {
		return nil, err
	}
	return resp.Data.Collections, nil
}

func (c Client) getCollectibles(owner string, collectionID string, chainId int) ([]Collectible, error) {
	query := url.Values{
		"user_addr":     {owner},
		"contract_addr": {collectionID},
		"chain_id":      {strconv.Itoa(chainId)},
	}

	var resp CollectibleResponse
	err := c.Get(&resp, "api/v1/assets", query)
	if err != nil {
		return nil, err
	}
	return resp.Data.Collectibles, err
}

func (c Client) fetchTokenURI(url string) (CollectionInfo, error) {
	var info CollectionInfo
	err := c.Get(&info, url, nil)
	return info, err
}
