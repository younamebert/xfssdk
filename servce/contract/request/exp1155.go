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

type Exp1155MintArgs struct {
	Recipient string   `json:"recipient"`
	Tokenurl  string   `json:"tokenurl"`
	Amount    *big.Int `json:"amount"`
}

type Exp1155MintBatchArgs struct {
	Recipient string `json:"recipient"`
	Amounts   string `json:"amounts"`
	TokenUrls string `json:"tokenurls"`
}

type Exp1155BalanceOfArgs struct {
	Account string   `json:"account"`
	ID      *big.Int `json:"id"`
}

type Exp1155IsApprovedForAllArgs struct {
	Account  string `json:"account"`
	Operator string `json:"operator"`
}

type Exp1155SetApprovalForAllArgs struct {
	Operator string `json:"operator"`
	Approved bool   `json:"approved"`
	PriKey   string `json:"prikey"`
}

type Exp1155SafeTransferFromArgs struct {
	From   string   `json:"from"`
	To     string   `json:"to"`
	Id     *big.Int `json:"id"`
	Amount *big.Int `json:"amount"`
	Prikey string   `json:"prikey"`
}

type Exp1155SafeBatchTransferFromArgs struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Ids     string `json:"ids"`
	Amounts string `json:"amounts"`
	Prikey  string `json:"prikey"`
}

type Exp1155BurnArgs struct {
	From   string   `json:"from"`
	ID     *big.Int `json:"id"`
	Amount *big.Int `json:"amount"`
	Prikey string   `json:"prikey"`
}

type Exp1155BurnBatchArgs struct {
	From    string `json:"from"`
	IDs     string `json:"ids"`
	Amounts string `json:"amounts"`
	Prikey  string `json:"prikey"`
}
