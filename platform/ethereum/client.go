package ethereum

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
)

type EthereumClient interface {
	GetTransactions(address string, coinIndex uint) (blockatlas.TxPage, error)
	GetTokenTxs(address, token string, coinIndex uint) (blockatlas.TxPage, error)
	GetTokenList(address string, coinIndex uint) (blockatlas.TokenPage, error)
	GetCurrentBlockNumber() (int64, error)
	GetBlockByNumber(num int64, coinIndex uint) (*blockatlas.Block, error)
}

type CollectibleClient interface {
	GetCollections(owner string, coinIndex uint) (blockatlas.CollectionPage, error)
	GetCollectibles(owner, collectionID string, coinIndex uint) (blockatlas.CollectiblePage, error)
}
