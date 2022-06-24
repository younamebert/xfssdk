package reqcontract

type NFTTokenCreateArgs struct {
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}

type DeployNFTokenArgs struct {
	Code       string `json:"code"`
	Addresskey string `json:"address_key"`
}

// mint request
type NFTokenMintArgs struct {
	// BlockChain      string `json:"blockchain"`
	// FromAddress     string `json:"from_address"`
	// ContractAddress string `json:"contract_address"`
	// MintAddress string `json:"mint_address"`
	TokenId string `json:"token_id"`
}

type NFTBalanceOfArgs struct {
	BalanceOfAddress string `json:"balance_of_address"`
}

type NFTOwnerOfArgs struct {
	TokenId string `json:"token_id"`
}

type NFTISApproveForAllArgs struct {
	OwnerAddress   string `json:"owner_address"`
	SpenderAddress string `json:"spender_address"`
}

type NFTGetApprovedArgs struct {
	TokenId string `json:"token_id"`
}

// type MintArgs struct {
// 	MintAddress string `json:"mint_address"`
// }

type NFTApproveArgs struct {
	ApproveFromAddressPriKey string `json:"approve_from_address_prikey"`
	ApproveToAddress         string `json:"approve_to_address"`
	Openid                   string `json:"openid"`
}

type NFTSetApproveForAllArgs struct {
	ApproveallFromAddressPriKey string `json:"approveall_from_adress_prikey"`
	ApproveallToAddress         string `json:"approveall_to_adress"`
	AllApproved                 bool   `json:"all_approved"`
}

type NFTTransferFromArgs struct {
	TransferOperatorAddressPriKey string `json:"transfer_operator_address_prikey"`
	TransferFromAddress           string `json:"transfer_from_address"`
	TransferToAddress             string `json:"transfer_to_address"`
	TransferFromTokenId           string `json:"transfer_from_tokenid"`
}
