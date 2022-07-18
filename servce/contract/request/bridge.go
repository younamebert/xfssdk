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
	TransferToAddress         string `json:"transfer_to_address"`
	TransferAmount            string `json:"transfer_amount"`
}
