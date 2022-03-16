package apitxpool

import (
	"xfssdk/core/apis"
)

type TxPoolLink interface {
	GetPending() (*apis.TransactionsResp, error)
	GetQueue() (*apis.TransactionsResp, error)
	GetPendingSize() (*int, error)
	GetTranByHash(hash string) (*apis.TransactionResp, error)
	GetAddrTxNonce(address string) (*int64, error)
	SendRawTransaction(data string) (*string, error)
}

type ApiTxPool struct {
	// XFSCLICENT *client.Client
}

// func NewApiTxPool(cli *client.Client) *ApiTxPool {
// 	return &ApiTxPool{
// 		XFSCLICENT: cli,
// 	}
// }

func (txpool *ApiTxPool) GetPending() (*apis.TransactionsResp, error) {

	result := new(apis.TransactionsResp)
	if err := apis.XFSCLICENT.CallMethod(1, "TxPool.GetPending", nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (txpool *ApiTxPool) GetQueue() (*apis.TransactionsResp, error) {
	result := new(apis.TransactionsResp)
	if err := apis.XFSCLICENT.CallMethod(1, "TxPool.GetQueue", nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (txpool *ApiTxPool) GetPendingSize() (*int, error) {
	var result *int
	if err := apis.XFSCLICENT.CallMethod(1, "TxPool.GetQueue", nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// func (tx *TxPoolHandler) Clear(_ EmptyArgs, resp *string) error {
// 	tx.TxPool.RemoveTransactions(tx.TxPool.GetTransactions())
// 	return nil
// }

func (txpool *ApiTxPool) GetTranByHash(hash string) (*apis.TransactionResp, error) {
	req := &GetTranByHashArgs{
		Hash: hash,
	}

	result := new(apis.TransactionResp)
	if err := apis.XFSCLICENT.CallMethod(1, "TxPool.GetTranByHash", &req, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (txpool *ApiTxPool) GetAddrTxNonce(address string) (*int64, error) {
	req := &GetAddrNonceByHashArgs{
		Address: address,
	}
	var result *int64
	if err := apis.XFSCLICENT.CallMethod(1, "TxPool.GetAddrTxNonce", &req, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (txpool *ApiTxPool) SendRawTransaction(data string) (*string, error) {
	req := &RawTransactionArgs{
		Data: data,
	}
	var result *string
	if err := apis.XFSCLICENT.CallMethod(1, "TxPool.SendRawTransaction", &req, &result); err != nil {
		return nil, err
	}
	return result, nil
}
