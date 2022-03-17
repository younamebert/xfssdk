package api

import (
	apichain "github.com/younamebert/xfssdk/api/chain"
	apinet "github.com/younamebert/xfssdk/api/net"
	apistate "github.com/younamebert/xfssdk/api/state"
	apitxpool "github.com/younamebert/xfssdk/api/txpool"
)

type ApiMethod struct {
	Chain  apichain.ChainLink
	Net    apinet.NetLink
	State  apistate.StateLink
	TxPool apitxpool.TxPoolLink
}

func NewApiMethod() *ApiMethod {
	return &ApiMethod{
		Chain:  new(apichain.ApiChain),
		Net:    new(apinet.ApiNet),
		State:  new(apistate.ApiState),
		TxPool: new(apitxpool.ApiTxPool),
	}
}
