package tezos

import (
	"github.com/trustwallet/blockatlas/pkg/blockatlas"
	"github.com/trustwallet/golibs/coin"
)

type Platform struct {
	client      Client
	rpcClient   RpcClient
	bakerClient BakerClient
}

func Init(api, rpc, baker string) *Platform {
	p := &Platform{
		client:      Client{blockatlas.InitClient(api)},
		rpcClient:   RpcClient{blockatlas.InitClient(rpc)},
		bakerClient: BakerClient{blockatlas.InitClient(baker)},
	}
	p.client.SetTimeout(35)
	return p
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[coin.XTZ]
}
