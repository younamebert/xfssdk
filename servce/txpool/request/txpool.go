package reqtxpool

type GetTranByHashArgs struct {
	Hash string `json:"hash"`
}

type GetAddrNonceByHashArgs struct {
	Address string `json:"address"`
}

type RemoveTxHashArgs struct {
	Hash string `json:"hash"`
}
