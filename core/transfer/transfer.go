package transfer

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/younamebert/xfssdk/core/apis"
	"github.com/younamebert/xfssdk/libs"
	reqtransfer "github.com/younamebert/xfssdk/servce/transfer/request"
	reqtxpool "github.com/younamebert/xfssdk/servce/txpool/request"
)

// EnCodeRawTransaction Create a signed transaction
func EnCodeRawTransaction(fromprikey string, tx *reqtransfer.StringRawTransaction) (*reqtransfer.StringRawTransaction, error) {
	if tx.Version == "" {
		tx.Version = "0"
	}

	if tx.Nonce == "" {
		address, err := libs.StrKey2Address(fromprikey)
		if err != nil {
			return nil, fmt.Errorf("invalid EnCodeRawTransaction fromprikey:%v", fromprikey)
		}
		reqGetNonce := &reqtxpool.GetAddrNonceByHashArgs{
			Address: address.B58String(),
		}
		var nonce *int64
		if err := apis.GVA_XFSCLICENT.CallMethod(1, "TxPool.GetAddrTxNonce", &reqGetNonce, &nonce); err != nil {
			return nil, fmt.Errorf("invalid GetAddrTxNonce addr:%v err:%v", address.B58String(), err)
		}
		tx.Nonce = strconv.FormatInt(*nonce, 10)
	}

	if err := tx.SignWithPrivateKey(fromprikey); err != nil {
		return nil, err
	}
	return tx, nil
}

// DeCodeRawTransaction Decode a signed transaction into a stringrawtransaction object
func DeCodeRawTransaction(dataraw string) (*reqtransfer.StringRawTransaction, error) {

	databytes, err := base64.StdEncoding.DecodeString(dataraw)
	if err != nil {
		return nil, err
	}

	rawtx := &reqtransfer.StringRawTransaction{}
	if err := json.Unmarshal(databytes, rawtx); err != nil {
		return nil, fmt.Errorf("failed to parse data: %s", err)
	}

	// if err := rawtx.SignWithPrivateKey(fromprikey); err != nil {
	// 	return nil, err
	// }
	return rawtx, nil
}

// SendRawTransfer Send a transaction to the specified node chain
func SendRawTransfer(basetransfer string) (string, error) {
	var txhash *string

	reqSendRawTransfer := &reqtransfer.RawTransactionArgs{
		Data: basetransfer,
	}
	if err := apis.GVA_XFSCLICENT.CallMethod(1, "TxPool.SendRawTransaction", &reqSendRawTransfer, &txhash); err != nil {
		return "", err
	}

	return *txhash, nil
}
