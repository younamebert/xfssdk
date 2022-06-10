package reqcontract

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

type StdTokenAllowanceArgs struct {
	OwnerAddress   string `json:"owner_address"`
	SpenderAddress string `json:"spender_address"`
}

// mint request
type StdTokenMintArgs struct {
	// BlockChain      string `json:"blockchain"`
	// FromAddress     string `json:"from_address"`
	// ContractAddress string `json:"contract_address"`
	MintAddress string `json:"mint_address"`
	Amount      string `json:"amount"`
}

// Burn requset
type StdTokenBurnArgs struct {
	BurnAddress string `json:"burn_address"`
	Amount      string `json:"amount"`
}

// Approve requset

type StdTokenApproveArgs struct {
	// BlockChain      string `json:"blockchain"`
	// FromAddress     string `json:"from_address"`
	// ContractAddress string `json:"contract_address"`
	ApproveFromAddressPriKey string `json:"approve_from_address_prikey"`
	ApproveSpenderAddress    string `json:"approve_spender_address"`
	Amount                   string `json:"amount"`
}

// TransferFrom requset
type StdTokenTransferFromArgs struct {
	// BlockChain  string `json:"blockchain"`
	TransferSpenderAddressPriKey string `json:"transfer_spender_address_prikey"`
	TransferFromAddress          string `json:"transfer_from_address"`
	TransferToAddress            string `json:"transfer_to_address"`
	TransferFromValue            string `json:"transfer_from_value"`
}
