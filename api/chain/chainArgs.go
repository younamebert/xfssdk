package apichain

type GetBlockByNumArgs struct {
	Number string `json:"number"`
}

type GetBlockByHashArgs struct {
	Hash string `json:"hash"`
}

type GetTxsByBlockNumArgs struct {
	Number string `json:"number"`
}

type GetTxbyBlockHashArgs struct {
	Hash string `json:"hash"`
}

// type GetBalanceOfAddressArgs struct {
// 	Address string `json:"address"`
// }

type GetTransactionArgs struct {
	Hash string `json:"hash"`
}

type GetReceiptByHashArgs struct {
	Hash string `json:"hash"`
}

type GetBlockHeaderByNumberArgs struct {
	Number string `json:"number"`
	//Count  string `json:"count"`
}

type GetBlockHeaderByHashArgs struct {
	Hash string `json:"hash"`
}

type GetBlocksByRangeArgs struct {
	From  string `json:"from"`
	Count string `json:"count"`
}

type GetBlocksArgs struct {
	Blocks string `json:"blocks"`
}

type ProgressBarArgs struct {
	Number int `json:"number"`
}

type GetBlockTxCountByHashArgs struct {
	Hash string `json:"hash"`
}

type GetBlockTxCountByNumArgs struct {
	Number string `json:"number"`
}

type GetBlockTxByHashAndIndexArgs struct {
	Hash  string `json:"hash"`
	Index int    `json:"index"`
}

type GetBlockTxByNumAndIndexArgs struct {
	Number string `json:"number"`
	Index  int    `json:"index"`
}
type GetBlockHashesArgs struct {
	Number string `json:"number"`
	Count  string `json:"count"`
}
