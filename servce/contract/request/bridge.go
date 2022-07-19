package reqcontract

import "math/big"

type BridgeArgs struct {
	Name            string   `json:"name"`
	Symbol          string   `json:"symbol"`
	ContractAddress string   `json:"ContractAddress"`
	ChainId         *big.Int `json:"chainId"`
}

type DeployBridgeArgs struct {
	Code       string `json:"code"`
	Addresskey string `json:"address_key"`
}

type BridgeTransferInArgs struct {
	TransferFromAddressPriKey string `json:"approve_from_address_prikey"`
	DepositorAddress          string `json:"depositor_address"`
	TransferAmount            string `json:"transfer_amount"`
	TransferFromChainId       string `json:"transfer_from_chainid"`
}

type BridgeTransferOutArgs struct {
	TransferFromAddressPriKey string `json:"approve_from_address_prikey"`
	TransferFromAddress       string `json:"transfer_From_address"` //可以是储户地址(或者是储户的授权地址)
	TransferToAddress         string `json:"transfer_to_address"`   //其他链地址
	TransferAmount            string `json:"transfer_amount"`
	TransferToChainId         string `json:"transfer_to_chainid"`
}
