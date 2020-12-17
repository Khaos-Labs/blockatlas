package ethereum

import "github.com/trustwallet/blockatlas/pkg/blockatlas"

func (p *Platform) GetCollections(owner string) (blockatlas.CollectionPage, error) {
	return p.collectible.GetCollections(owner, p.CoinIndex)
}

func (p *Platform) GetCollectibles(owner, collectionID string) (blockatlas.CollectiblePage, error) {
	return p.collectible.GetCollectibles(owner, collectionID, p.CoinIndex)
}
