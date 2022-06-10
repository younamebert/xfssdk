package resptransfer

import (
	"math/big"

	"github.com/younamebert/xfssdk/common"
)

type TransactionResp struct {
	Version   uint32         `json:"version"`
	To        common.Address `json:"to"`
	GasPrice  *big.Int       `json:"gas_price"`
	GasLimit  *big.Int       `json:"gas_limit"`
	Nonce     uint64         `json:"nonce"`
	Value     *big.Int       `json:"value"`
	From      string         `json:"from"`
	Hash      common.Hash    `json:"hash"`
	Data      []byte         `json:"data"`
	Signature []byte         `json:"signature"`
}

type TransactionsResp []*TransactionResp

type ReceiptResp struct {
	Version     uint32      `json:"version"`
	Status      uint32      `json:"status"`
	TxHash      common.Hash `json:"tx_hash"`
	GasUsed     *big.Int    `json:"gas_used"`
	BlockHeight uint64      `json:"block_height"`
	BlockHash   common.Hash `json:"block_hash"`
	BlockIndex  uint64      `json:"block_index"`
	TxIndex     uint64      `json:"tx_index"`
}
