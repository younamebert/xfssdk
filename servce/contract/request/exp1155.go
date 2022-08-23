package reqcontract

import (
	"math/big"
)

type Exp1155CreateArgs struct {
}

type Exp1155DeployArgs struct {
	Privkey string `json:"privkey"`
	Code    string `json:"code"`
}

// type Exp1155MintArgs struct {
// 	Address  string   `json:"address"`
// 	Amount   *big.Int `json:"amount"`
// 	TokenUrl string   `json:"token_uri"`
// }

type Exp1155MintArgs struct {
	Recipient string `json:"recipient"`
	Tokenurl  string `json:"tokenurls"`
}

type Exp1155MintBatchArgs struct {
	Address   string   `json:"address"`
	Amounts   []string `json:"amount"`
	TokenUrls []string `json:"token_url"`
}

type Exp1155BalanceOfArgs struct {
	Account string   `json:"account"`
	ID      *big.Int `json:"id"`
}

type Exp1155BalanceOfBatchArgs struct {
	Accounts []string   `json:"accounts"`
	IDS      []*big.Int `json:"ids"`
}

type Exp1155IsApprovedForAllArgs struct {
	Account  string `json:"account"`
	Operator string `json:"operator"`
}

type Exp1155SetApprovalForAllArgs struct {
	Operator string `json:"operator"`
	Approved bool   `json:"approved"`
}

type Exp1155SafeTransferFromArgs struct {
	From   string   `json:"from"`
	To     string   `json:"to"`
	Id     string   `json:"id"`
	Amount *big.Int `json:"amount"`
}

type Exp1155SafeBatchTransferFromArgs struct {
	From    string     `json:"from"`
	To      string     `json:"to"`
	Ids     []*big.Int `json:"ids"`
	Amounts []*big.Int `json:"amounts"`
}

type Exp1155BurnArgs struct {
	From   string   `json:"from"`
	ID     *big.Int `json:"id"`
	Amount *big.Int `json:"amount"`
}

type Exp1155BurnBatchArgs struct {
	From    string     `json:"from"`
	IDs     []*big.Int `json:"ids"`
	Amounts []*big.Int `json:"amounts"`
}
