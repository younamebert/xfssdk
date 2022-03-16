package api

import (
	apichain "xfssdk/api/chain"
	apinet "xfssdk/api/net"
	apistate "xfssdk/api/state"
	apitxpool "xfssdk/api/txpool"
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
