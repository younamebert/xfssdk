package api

import (
	apichain "github.com/younamebert/xfssdk/api/chain"
	apicontract "github.com/younamebert/xfssdk/api/contract"
	apinet "github.com/younamebert/xfssdk/api/net"
	apistate "github.com/younamebert/xfssdk/api/state"
	apitransfer "github.com/younamebert/xfssdk/api/transfer"
	apitxpool "github.com/younamebert/xfssdk/api/txpool"
)

// ApiMethod node chain method set structure
type ApiMethod struct {
	Chain    apichain.ChainLink
	Net      apinet.NetLink
	State    apistate.StateLink
	TxPool   apitxpool.TxPoolLink
	Contract apicontract.TokenLink
	Transfer apitransfer.TransferLink
}

func NewApiMethod() *ApiMethod {
	return &ApiMethod{
		Chain:    new(apichain.ApiChain),
		Net:      new(apinet.ApiNet),
		State:    new(apistate.ApiState),
		TxPool:   new(apitxpool.ApiTxPool),
		Contract: new(apicontract.ApiToken),
		Transfer: new(apitransfer.ApiTransfer),
	}
}
