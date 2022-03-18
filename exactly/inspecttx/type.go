package inspecttx

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"sort"

	"github.com/younamebert/xfssdk/common"
	"github.com/younamebert/xfssdk/libs/ahash"
	"github.com/younamebert/xfssdk/libs/crypto"
)

// type Transaction struct {
// 	Version   uint32         `json:"version"`
// 	From      common.Address `json:"from"`
// 	To        common.Address `json:"to"`
// 	GasPrice  *big.Int       `json:"gas_price"`
// 	GasLimit  *big.Int       `json:"gas_limit"`
// 	Data      []byte         `json:"data"`
// 	Nonce     uint64         `json:"nonce"`
// 	Value     *big.Int       `json:"value"`
// 	Signature []byte         `json:"signature"`
// }

type StringRawTransaction struct {
	Version   string `json:"version"`
	To        string `json:"to"`
	Value     string `json:"value"`
	Data      string `json:"data"`
	GasLimit  string `json:"gas_limit"`
	GasPrice  string `json:"gas_price"`
	Signature string `json:"signature"`
	Nonce     string `json:"nonce"`
}

// SignWithPrivateKey Generate a transaction signature
func (tx *StringRawTransaction) SignWithPrivateKey(fromprikey string) error {

	keyEnc := fromprikey
	if keyEnc[0] == '0' && keyEnc[1] == 'x' {
		keyEnc = keyEnc[2:]
	} else {
		return errors.New("binary forward backward error")
	}

	keyDer, err := hex.DecodeString(keyEnc)
	if err != nil {
		return err
	}

	_, key, err := crypto.DecodePrivateKey(keyDer)
	if err != nil {
		return err
	}

	hash := tx.SignHash()
	sig, err := crypto.ECDSASign(hash.Bytes(), key)
	if err != nil {
		return err
	}
	tx.Signature = hex.EncodeToString(sig)
	return nil
}

// RawTx trading partner code Base64 format
func (tx *StringRawTransaction) RawTx() (string, error) {
	bs, err := json.Marshal(tx)
	if err != nil {
		return "", err
	}
	result := base64.StdEncoding.EncodeToString(bs)
	return result, nil
}

// signHash generate transaction hash
func (tx *StringRawTransaction) SignHash() common.Hash {
	//nt := t.copyTrim()

	tmp := map[string]string{
		"version":   tx.Value,
		"to":        tx.To,
		"gas_price": tx.GasPrice,
		"gas_limit": tx.GasLimit,
		"data":      tx.Data,
		"nonce":     tx.Nonce,
		"value":     tx.Value,
	}
	enc := sortAndEncodeMap(tmp)
	if enc == "" {
		return common.Hash{}
	}
	return common.Bytes2Hash(ahash.SHA256([]byte(enc)))
}

func sortAndEncodeMap(data map[string]string) string {
	mapkeys := make([]string, 0)
	for k := range data {
		mapkeys = append(mapkeys, k)
	}
	sort.Strings(mapkeys)
	strbuf := ""
	for i, key := range mapkeys {
		val := data[key]
		if val == "" {
			continue
		}
		strbuf += fmt.Sprintf("%s=%s", key, val)
		if i < len(mapkeys)-1 {
			strbuf += "&"
		}
	}
	return strbuf
}
