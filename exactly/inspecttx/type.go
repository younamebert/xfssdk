package inspecttx

import (
	"fmt"
	"math/big"
	"sort"
	"strconv"
	"xfssdk/common"
	"xfssdk/libs/crypto"
)

type Transaction struct {
	Version   uint32         `json:"version"`
	From      common.Address `json:"from"`
	To        common.Address `json:"to"`
	GasPrice  *big.Int       `json:"gas_price"`
	GasLimit  *big.Int       `json:"gas_limit"`
	Data      []byte         `json:"data"`
	Nonce     uint64         `json:"nonce"`
	Value     *big.Int       `json:"value"`
	Signature []byte         `json:"signature"`
}

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

func coverTransaction(r *StringRawTransaction) (*Transaction, error) {
	version, err := strconv.ParseInt(r.Version, 10, 32)
	if err != nil {
		return nil, fmt.Errorf("failed to parse version: %s", err)
	}
	toaddr := common.ZeroAddr
	if r.To != "" {
		toaddr = common.StrB58ToAddress(r.To)
		if !crypto.VerifyAddress(toaddr) {
			return nil, fmt.Errorf("failed to verify 'to' address: %s", r.To)
		}
	} else if r.Data == "" {
		return nil, fmt.Errorf("failed to parse 'to' address")
	}
	gasprice, ok := new(big.Int).SetString(r.GasPrice, 10)
	if !ok {
		return nil, fmt.Errorf("failed to parse gasprice")
	}
	gaslimit, ok := new(big.Int).SetString(r.GasLimit, 10)
	if !ok {
		return nil, fmt.Errorf("failed to parse gasprice")
	}
	data := common.Hex2bytes(r.Data)
	nonce, err := strconv.ParseInt(r.Nonce, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse nonce: %s", err)
	}
	value, ok := new(big.Int).SetString(r.Value, 10)
	if !ok {
		return nil, fmt.Errorf("failed to parse value")
	}
	return &Transaction{
		Version:  uint32(version),
		To:       toaddr,
		GasPrice: gasprice,
		GasLimit: gaslimit,
		Data:     data,
		Nonce:    uint64(nonce),
		Value:    value,
	}, nil
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
