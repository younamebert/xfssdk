package inspecttx

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/younamebert/xfssdk/common"
	"github.com/younamebert/xfssdk/libs/crypto"
)

type InspectTxWay interface {
	NewRawTransaction(fromprikey string, tx *StringRawTransaction) (*StringRawTransaction, error)
	GetFromAddress(fromprikey string) (common.Address, error)
	CoverTransaction(fromprikey, dataraw string) (*StringRawTransaction, error)
}

type InspectTx struct{}

//GetFromAddress 根据from的私钥获取地址
func (inspecttx *InspectTx) GetFromAddress(fromprikey string) (common.Address, error) {

	keyEnc := fromprikey

	if keyEnc[0] == '0' && keyEnc[1] == 'x' {
		keyEnc = keyEnc[2:]
	} else {
		return common.Address{}, errors.New("binary forward backward error")
	}

	keyDer, err := hex.DecodeString(keyEnc)
	if err != nil {
		return common.Address{}, err
	}

	_, pKey, err := crypto.DecodePrivateKey(keyDer)
	if err != nil {
		return common.Address{}, err
	}

	addr := crypto.DefaultPubKey2Addr(pKey.PublicKey)
	return addr, nil
}

//NewRawTransaction 创建一个签名交易
func (inspecttx *InspectTx) NewRawTransaction(fromprikey string, tx *StringRawTransaction) (*StringRawTransaction, error) {

	if err := tx.SignWithPrivateKey(fromprikey); err != nil {
		return nil, err
	}
	return tx, nil
}

//CoverTransaction 解码一个签名交易转tx对象
func (inspecttx *InspectTx) CoverTransaction(fromprikey, dataraw string) (*StringRawTransaction, error) {

	databytes, err := base64.StdEncoding.DecodeString(dataraw)
	if err != nil {
		return nil, err
	}

	rawtx := &StringRawTransaction{}
	if err := json.Unmarshal(databytes, rawtx); err != nil {
		return nil, fmt.Errorf("failed to parse data: %s", err)
	}

	if err := rawtx.SignWithPrivateKey(fromprikey); err != nil {
		return nil, err
	}
	return rawtx, nil
}
