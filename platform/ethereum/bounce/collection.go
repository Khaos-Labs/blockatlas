package bounce

import (
	"strconv"

	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

var (
	bscChainId = 56
)

func (c *Client) GetCollections(owner string, coinIndex uint) (blockatlas.CollectionPage, error) {
	collections, err := c.getCollections(owner, bscChainId)
	if err != nil {
		return nil, err
	}
	return c.NormalizeCollections(collections, coinIndex, owner)

}

func (c *Client) GetCollectibles(owner, collectionID string, coinIndex uint) (blockatlas.CollectiblePage, error) {
	collectibles, err := c.getCollectibles(owner, collectionID, bscChainId)
	if err != nil {
		return nil, err
	}
	return c.NormalizeCollectibles(collectibles, coinIndex)
}

func (c *Client) NormalizeCollections(collections []Collection, coinIndex uint, owner string) (blockatlas.CollectionPage, error) {
	page := make(blockatlas.CollectionPage, len(collections))
	for _, cl := range collections {
		total, err := strconv.Atoi(cl.Balance)
		if err != nil {
			continue
		}
		// skip empty info token
		if len(cl.TokenURI) == 0 {
			continue
		}
		info, err := c.fetchTokenURI(cl.TokenURI)
		if err != nil {
			return nil, err
		}
		page = append(page, blockatlas.Collection{
			Id:           cl.ContractAddr,
			Name:         info.Properties.Name.Description,
			ImageUrl:     info.Properties.Image.Description,
			Description:  info.Properties.Description.Description,
			ExternalLink: cl.TokenURI,
			Total:        total,
			Address:      owner,
			Coin:         coinIndex,
			Type:         "ERC" + cl.TokenType,
		})
	}
	return page, nil
}

func (c *Client) NormalizeCollectibles(collectibles []Collectible, coinIndex uint) (blockatlas.CollectiblePage, error) {
	if len(collectibles) == 0 {
		return blockatlas.CollectiblePage{}, nil
	}
	page := make(blockatlas.CollectiblePage, len(collectibles))
	info, err := c.fetchTokenURI(collectibles[0].TokenURI)
	if err != nil {
		return nil, err
	}
	for _, c := range collectibles {
		page = append(page, blockatlas.Collectible{
			ID:              genId(c.ID),
			CollectionID:    c.ContractAddr,
			TokenID:         strconv.Itoa(c.TokenID),
			ContractAddress: c.ContractAddr,
			ImageUrl:        info.Properties.Image.Description,
			ExternalLink:    c.TokenURI,
			Type:            "ERC721",
			Description:     info.Properties.Description.Description,
			Coin:            coinIndex,
			Name:            info.Properties.Name.Description,
			Version:         "3.0",
		})
	}
	return page, nil
}

func genId(id int) string {
	return "bounce-" + strconv.Itoa(id)
}
