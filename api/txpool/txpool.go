package apitxpool

import (
	"github.com/younamebert/xfssdk/core/apis"
)

type TxPoolLink interface {
	GetPending() (*apis.TransactionsResp, error)
	GetQueue() (*apis.TransactionsResp, error)
	GetPendingSize() (*int, error)
	GetTranByHash(hash string) (*apis.TransactionResp, error)
	GetAddrTxNonce(address string) (*int64, error)
	SendRawTransaction(data string) (*string, error)
}

type ApiTxPool struct{}

// GetPending 获取交易池的pending队列所有的交易信息
func (txpool *ApiTxPool) GetPending() (*apis.TransactionsResp, error) {

	result := new(apis.TransactionsResp)
	if err := apis.GVA_XFSCLICENT.CallMethod(1, "TxPool.GetPending", nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetQueue 获取交易池的queue队列所有的交易信息
func (txpool *ApiTxPool) GetQueue() (*apis.TransactionsResp, error) {
	result := new(apis.TransactionsResp)
	if err := apis.GVA_XFSCLICENT.CallMethod(1, "TxPool.GetQueue", nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetPendingSize 获取当前交易池的可以执行交易数量
func (txpool *ApiTxPool) GetPendingSize() (*int, error) {
	var result *int
	if err := apis.GVA_XFSCLICENT.CallMethod(1, "TxPool.GetQueue", nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetTranByHash 指定交易txhash在交易池获取交易信息
func (txpool *ApiTxPool) GetTranByHash(hash string) (*apis.TransactionResp, error) {
	req := &GetTranByHashArgs{
		Hash: hash,
	}

	result := new(apis.TransactionResp)
	if err := apis.GVA_XFSCLICENT.CallMethod(1, "TxPool.GetTranByHash", &req, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (txpool *ApiTxPool) GetAddrTxNonce(address string) (*int64, error) {
	req := &GetAddrNonceByHashArgs{
		Address: address,
	}
	var result *int64
	if err := apis.GVA_XFSCLICENT.CallMethod(1, "TxPool.GetAddrTxNonce", &req, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (txpool *ApiTxPool) SendRawTransaction(data string) (*string, error) {
	req := &RawTransactionArgs{
		Data: data,
	}
	var result *string
	if err := apis.GVA_XFSCLICENT.CallMethod(1, "TxPool.SendRawTransaction", &req, &result); err != nil {
		return nil, err
	}
	return result, nil
}
