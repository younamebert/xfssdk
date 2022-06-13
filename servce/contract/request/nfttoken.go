package reqcontract

type NFTTokenCreateArgs struct {
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}

type BalanceOfArgs struct {
	BalanceOfAddress string `json:"balance_of_address"`
}

type OwnerOfArgs struct {
	TokenId string `json:"token_id"`
}

type ISApproveForAllArgs struct {
	OwnerAddress   string `json:"owner_address"`
	SpenderAddress string `json:"spender_address"`
}

type GetApprovedArgs struct {
	TokenId string `json:"token_id"`
}

type MintArgs struct {
	MintAddress string `json:"mint_address"`
}

type ApproveArgs struct {
	ApproveFromAddressPriKey string `json:"approve_from_address_prikey"`
	ApproveToAddress         string `json:"approve_to_address"`
	Openid                   string `json:"openid"`
}

type SetApproveForAllArgs struct {
	ApproveallFromAddressPriKey string `json:"approveall_from_adress_prikey"`
	ApproveallToAddress         string `json:"approveall_to_adress"`
	AllApproved                 bool   `json:"all_approved"`
}

type TransferFromArgs struct {
	TransferOperatorAddressPriKey string `json:"transfer_operator_address_prikey"`
	TransferFromAddress           string `json:"transfer_from_address"`
	TransferToAddress             string `json:"transfer_to_address"`
	TransferFromTokenId           string `json:"transfer_from_tokenid"`
}
