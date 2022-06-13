package apitransfer

import (
	"encoding/base64"
	"encoding/json"
	"fmt"

	reqtransfer "github.com/younamebert/xfssdk/servce/transfer/request"
)

type TransferLink interface {
	EnCodeRawTransaction(fromprikey string, tx *reqtransfer.StringRawTransaction) (string, error)
	DeCodeRawTransaction(dataraw string) (*reqtransfer.StringRawTransaction, error)
}

type ApiTransfer struct{}

// EnCodeRawTransaction Create a signed transaction
func (transfer *ApiTransfer) EnCodeRawTransaction(fromprikey string, tx *reqtransfer.StringRawTransaction) (string, error) {

	if err := tx.SignWithPrivateKey(fromprikey); err != nil {
		return "", err
	}
	return tx.Transfer2Raw()
}

// DeCodeRawTransaction Decode a signed transaction into a stringrawtransaction object
func (transfer *ApiTransfer) DeCodeRawTransaction(dataraw string) (*reqtransfer.StringRawTransaction, error) {

	databytes, err := base64.StdEncoding.DecodeString(dataraw)
	if err != nil {
		return nil, err
	}

	rawtx := &reqtransfer.StringRawTransaction{}
	if err := json.Unmarshal(databytes, rawtx); err != nil {
		return nil, fmt.Errorf("failed to parse data: %s", err)
	}
	return rawtx, nil
}
