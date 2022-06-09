package apitransfer

import (
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/younamebert/xfssdk/common"
	servetxpool "github.com/younamebert/xfssdk/servce/txpool"
)

type TransferLink interface {
	EnCodeRawTransaction(fromprikey string, tx *servetxpool.StringRawTransaction) (*servetxpool.StringRawTransaction, error)
	DeCodeRawTransaction(dataraw string) (*servetxpool.StringRawTransaction, error)
	GetFromAddress(fromprikey string) (common.Address, error)
	CheckWalletPriKey(der string) error
}

type ApiTransfer struct{}

// CheckWalletPriKey Verify that a private key complies with the rules
func (transfer *ApiTransfer) CheckWalletPriKey(der string) error {
	return common.CheckWalletPriKey(der)
	// servetxpool.StringRawTransaction
}

// GetFromAddress get address through private key
func (transfer *ApiTransfer) GetFromAddress(fromprikey string) (common.Address, error) {
	return common.StrKey2Address(fromprikey)
}

// EnCodeRawTransaction Create a signed transaction
func (transfer *ApiTransfer) EnCodeRawTransaction(fromprikey string, tx *servetxpool.StringRawTransaction) (string, error) {

	if err := tx.SignWithPrivateKey(fromprikey); err != nil {
		return "", err
	}
	return tx.Transfer2Raw()
}

// DeCodeRawTransaction Decode a signed transaction into a stringrawtransaction object
func (transfer *ApiTransfer) DeCodeRawTransaction(dataraw string) (*servetxpool.StringRawTransaction, error) {

	databytes, err := base64.StdEncoding.DecodeString(dataraw)
	if err != nil {
		return nil, err
	}

	rawtx := &servetxpool.StringRawTransaction{}
	if err := json.Unmarshal(databytes, rawtx); err != nil {
		return nil, fmt.Errorf("failed to parse data: %s", err)
	}

	// if err := rawtx.SignWithPrivateKey(fromprikey); err != nil {
	// 	return nil, err
	// }
	return rawtx, nil
}
