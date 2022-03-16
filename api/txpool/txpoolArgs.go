package apitxpool

type GetTranByHashArgs struct {
	Hash string `json:"hash"`
}

type GetAddrNonceByHashArgs struct {
	Address string `json:"address"`
}

type RawTransactionArgs struct {
	Data string `json:"data"`
}

type RemoveTxHashArgs struct {
	Hash string `json:"hash"`
}
