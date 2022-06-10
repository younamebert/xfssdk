package api

import (
	apichain "github.com/younamebert/xfssdk/api/chain"
	apicontract "github.com/younamebert/xfssdk/api/contract"
	apinet "github.com/younamebert/xfssdk/api/net"
	apistate "github.com/younamebert/xfssdk/api/state"
	apitxpool "github.com/younamebert/xfssdk/api/txpool"
)

// ApiMethod node chain method set structure
type ApiMethod struct {
	Chain       apichain.ChainLink
	Net         apinet.NetLink
	State       apistate.StateLink
	TxPool      apitxpool.TxPoolLink
	BasicsToken apicontract.BasicsTokenLink
	StdToken    apichain.ChainLink
}

func NewApiMethod() *ApiMethod {
	return &ApiMethod{
		Chain:       new(apichain.ApiChain),
		Net:         new(apinet.ApiNet),
		State:       new(apistate.ApiState),
		TxPool:      new(apitxpool.ApiTxPool),
		BasicsToken: new(apicontract.ApiBasicsToken),
	}
}
