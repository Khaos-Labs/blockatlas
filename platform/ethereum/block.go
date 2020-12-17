package ethereum

import "github.com/trustwallet/blockatlas/pkg/blockatlas"

func (p *Platform) CurrentBlockNumber() (int64, error) {
	return p.client.GetCurrentBlockNumber()
}

func (p *Platform) GetBlockByNumber(num int64) (*blockatlas.Block, error) {
	return p.client.GetBlockByNumber(num, p.CoinIndex)
}
