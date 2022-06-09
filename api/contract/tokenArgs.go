package apicontract

type TokenArgs struct {
	Name        string `json:"name"`
	Symbol      string `json:"symbol"`
	Decimals    string `json:"decimals"`
	TotalSupply string `json:"totalSupply"`
}

type DeployTokenArgs struct {
	Code       string `json:"code"`
	Addresskey string `json:"address_key"`
}

type GetAddrNonceByHashArgs struct {
	Address string `json:"address"`
}
