package apitxpool

import (
	"github.com/younamebert/xfssdk/core/apis"
	reqtransfer "github.com/younamebert/xfssdk/servce/transfer/request"
	resptransfer "github.com/younamebert/xfssdk/servce/transfer/response"
	reqtxpool "github.com/younamebert/xfssdk/servce/txpool/request"
)

type TxPoolLink interface {
	GetPending() (*resptransfer.TransactionsResp, error)
	GetQueue() (*resptransfer.TransactionsResp, error)
	GetPendingSize() (*int, error)
	GetTranByHash(hash string) (*resptransfer.TransactionResp, error)
	GetAddrTxNonce(address string) (*int64, error)
	SendRawTransaction(data string) (*string, error)
}

type ApiTxPool struct{}

// GetPending get all transaction information of the pending queue of the transaction pool
func (txpool *ApiTxPool) GetPending() (*resptransfer.TransactionsResp, error) {

	result := new(resptransfer.TransactionsResp)
	if err := apis.GVA_XFSCLICENT.CallMethod(1, "TxPool.GetPending", nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetQueue obtain all transaction information of the queue queue of the transaction pool
func (txpool *ApiTxPool) GetQueue() (*resptransfer.TransactionsResp, error) {
	result := new(resptransfer.TransactionsResp)
	if err := apis.GVA_XFSCLICENT.CallMethod(1, "TxPool.GetQueue", nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetPendingSize get the number of executable transactions in the current transaction pool
func (txpool *ApiTxPool) GetPendingSize() (*int, error) {
	var result *int
	if err := apis.GVA_XFSCLICENT.CallMethod(1, "TxPool.GetQueue", nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetTranByHash specify the transaction txhash to obtain transaction information in the transaction pool
func (txpool *ApiTxPool) GetTranByHash(hash string) (*resptransfer.TransactionResp, error) {
	req := &reqtxpool.GetTranByHashArgs{
		Hash: hash,
	}

	result := new(resptransfer.TransactionResp)
	if err := apis.GVA_XFSCLICENT.CallMethod(1, "TxPool.GetTranByHash", &req, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetAddrTxNonce specify the account address to obtain the account nonce value
func (txpool *ApiTxPool) GetAddrTxNonce(address string) (*int64, error) {
	req := &reqtxpool.GetAddrNonceByHashArgs{
		Address: address,
	}
	var result *int64
	if err := apis.GVA_XFSCLICENT.CallMethod(1, "TxPool.GetAddrTxNonce", &req, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// SendRawTransaction base64 send a transaction
func (txpool *ApiTxPool) SendRawTransaction(data string) (*string, error) {
	req := &reqtransfer.RawTransactionArgs{
		Data: data,
	}
	var result *string
	if err := apis.GVA_XFSCLICENT.CallMethod(1, "TxPool.SendRawTransaction", &req, &result); err != nil {
		return nil, err
	}
	return result, nil
}
