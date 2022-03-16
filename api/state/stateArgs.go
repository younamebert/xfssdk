package apistate

type GetAccountArgs struct {
	RootHash string `json:"root_hash"`
	Address  string `json:"address"`
}

type GetBalanceArgs struct {
	RootHash string `json:"root_hash"`
	Address  string `json:"address"`
}
