package respchain

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"math/big"

	"github.com/younamebert/xfssdk/common"
	"github.com/younamebert/xfssdk/common/ahash"
	"github.com/younamebert/xfssdk/crypto"
	resptransfer "github.com/younamebert/xfssdk/servce/transfer/response"
)

type Hashes []common.Hash

type BlockHeaderResp struct {
	Height        uint64         `json:"height"`
	Version       uint32         `json:"version"`
	HashPrevBlock common.Hash    `json:"hash_prev_block"`
	Timestamp     uint64         `json:"timestamp"`
	Coinbase      common.Address `json:"coinbase"`

	StateRoot        common.Hash `json:"state_root"`
	TransactionsRoot common.Hash `json:"transactions_root"`
	ReceiptsRoot     common.Hash `json:"receipts_root"`
	GasLimit         *big.Int    `json:"gas_limit"`
	GasUsed          *big.Int    `json:"gas_used"`

	Bits       uint32      `json:"bits"`
	Nonce      uint32      `json:"nonce"`
	ExtraNonce uint64      `json:"extranonce"`
	Hash       common.Hash `json:"hash"`
}

type BlockResp struct {
	Height        uint64         `json:"height"`
	Version       uint32         `json:"version"`
	HashPrevBlock common.Hash    `json:"hash_prev_block"`
	Timestamp     uint64         `json:"timestamp"`
	Coinbase      common.Address `json:"coinbase"`

	StateRoot        common.Hash `json:"state_root"`
	TransactionsRoot common.Hash `json:"transactions_root"`
	ReceiptsRoot     common.Hash `json:"receipts_root"`
	GasLimit         *big.Int    `json:"gas_limit"`
	GasUsed          *big.Int    `json:"gas_used"`

	Bits         uint32                        `json:"bits"`
	Nonce        uint32                        `json:"nonce"`
	ExtraNonce   uint64                        `json:"extranonce"`
	Hash         common.Hash                   `json:"hash"`
	Transactions resptransfer.TransactionsResp `json:"transactions"`
}

type RawTransactionArgs struct {
	Data string `json:"data"`
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

// Transfer2Raw trading partner code Base64 format
func (tx *StringRawTransaction) Transfer2Raw() (string, error) {
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
	enc := common.SortAndEncodeMap(tmp)
	if enc == "" {
		return common.Hash{}
	}
	return common.Bytes2Hash(ahash.SHA256([]byte(enc)))
}
