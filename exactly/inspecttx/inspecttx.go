package inspecttx

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/younamebert/xfssdk/common"
	"github.com/younamebert/xfssdk/libs/ahash"
)

type InspectTxWay interface {
	SignHashByData(dataraw string) (common.Hash, error)
	SignHash(tx *Transaction) (common.Hash, error)
}

type InspectTx struct {
}

func (inspecttx *InspectTx) SignHashByData(dataraw string) (common.Hash, error) {
	databytes, err := base64.StdEncoding.DecodeString(dataraw)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to parse data: %s", err)
	}
	rawtx := &StringRawTransaction{}
	if err := json.Unmarshal(databytes, rawtx); err != nil {
		return common.Hash{}, fmt.Errorf("failed to parse data: %s", err)
	}
	tx, err := coverTransaction(rawtx)
	if err != nil {
		return common.Hash{}, err
	}

	data := ""
	if tx.Data != nil && len(tx.Data) > 0 {
		data = "0x" + hex.EncodeToString([]byte(tx.Data))
	}
	tmp := map[string]string{
		"version":   strconv.FormatInt(int64(tx.Version), 10),
		"to":        tx.To.String(),
		"gas_price": tx.GasPrice.Text(10),
		"gas_limit": tx.GasLimit.Text(10),
		"data":      data,
		"nonce":     strconv.Itoa(int(tx.Nonce)),
		"value":     tx.Value.Text(10),
	}
	enc := sortAndEncodeMap(tmp)
	if enc == "" {
		return common.Hash{}, fmt.Errorf("sortAndEncodeMap err")
	}
	return common.Bytes2Hash(ahash.SHA256([]byte(enc))), nil
}

func (inspecttx *InspectTx) SignHash(tx *Transaction) (common.Hash, error) {

	if tx.To.Equals(common.Address{}) {
		return common.Hash{}, fmt.Errorf("transfer to not nil")
	}
	if tx.From.Equals(common.Address{}) {
		return common.Hash{}, fmt.Errorf("transfer from not nil")
	}

	data := ""
	if tx.Data != nil && len(tx.Data) > 0 {
		data = "0x" + hex.EncodeToString([]byte(tx.Data))
	}
	tmp := map[string]string{
		"version":   strconv.FormatInt(int64(tx.Version), 10),
		"to":        tx.To.String(),
		"gas_price": tx.GasPrice.Text(10),
		"gas_limit": tx.GasLimit.Text(10),
		"data":      data,
		"nonce":     strconv.Itoa(int(tx.Nonce)),
		"value":     tx.Value.Text(10),
	}
	enc := sortAndEncodeMap(tmp)
	if enc == "" {
		return common.Hash{}, fmt.Errorf("sortAndEncodeMap err")
	}
	return common.Bytes2Hash(ahash.SHA256([]byte(enc))), nil
}
