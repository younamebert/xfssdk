package apis

import (
	"math/big"

	"github.com/younamebert/xfssdk/common"
)

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

	Bits         uint32           `json:"bits"`
	Nonce        uint32           `json:"nonce"`
	ExtraNonce   uint64           `json:"extranonce"`
	Hash         common.Hash      `json:"hash"`
	Transactions TransactionsResp `json:"transactions"`
}

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

type Hashes []common.Hash

type StateObjResp struct {
	Balance   *string      `json:"balance"`
	Nonce     uint64       `json:"nonce"`
	Extra     *string      `json:"extra"`
	Code      *string      `json:"code"`
	StateRoot *common.Hash `json:"state_root"`
}
