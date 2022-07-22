package reqcontract

import "math/big"

type NftMarketArgs struct{}

type DeployNftMarketArgs struct {
	Code       string `json:"code"`
	Addresskey string `json:"address_key"`
}

type NFTMarketMintArgs struct {
	Address  string   `json:"address"`
	Amount   *big.Int `json:"amount"`
	TokenUrl string   `json:"token_uri"`
}

type NFTMarketMintBatchArgs struct {
	Address   string     `json:"address"`
	Amounts   []*big.Int `json:"amount"`
	TokenUrls []string   `json:"token_uri"`
}
