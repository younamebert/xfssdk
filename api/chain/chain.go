package apichain

import (
	"github.com/younamebert/xfssdk/common"
	"github.com/younamebert/xfssdk/core/apis"
	reqchain "github.com/younamebert/xfssdk/servce/chain/request"
	respchain "github.com/younamebert/xfssdk/servce/chain/response"
	resptransfer "github.com/younamebert/xfssdk/servce/transfer/response"
)

type ChainLink interface {
	GetBlockByNumber(number string) (*respchain.BlockResp, error)
	GetBlockHashes(startHeight, endHeight string) (*respchain.Hashes, error)
	GetHead() (*respchain.BlockHeaderResp, error)
	GetBlockHeaderByNumber(number string) (*respchain.BlockHeaderResp, error)
	GetBlockHeaderByHash(hash string) (*respchain.BlockHeaderResp, error)
	GetBlockByHash(hash string) (*respchain.BlockResp, error)
	GetTxsByBlockNum(number string) (*resptransfer.TransactionsResp, error)
	GetTxsByBlockHash(hash string) (*resptransfer.TransactionsResp, error)
	GetReceiptByHash(txhash string) (*resptransfer.ReceiptResp, error)
	GetTransaction(txhash string) (*resptransfer.TransactionResp, error)
	GetChainTransfer(txhash string) bool
	GetBlockTxCountByHash(hash string) (*int, error)
	GetBlockTxCountByNum(number string) (*int, error)
	GetBlockTxByHashAndIndex(hash string, index int) (*resptransfer.TransactionResp, error)
	GetBlockTxByNumAndIndex(number string, index int) (*resptransfer.TransactionResp, error)
}

type ApiChain struct{}

// GetBlockByNumber obtain block information according to block height
func (chain *ApiChain) GetBlockByNumber(number string) (*respchain.BlockResp, error) {

	if err := common.Str2Int64(number); err != nil {
		return nil, err
	}
	req := &reqchain.GetBlockByNumArgs{
		Number: number,
	}
	result := new(respchain.BlockResp)
	if err := apis.GVA_XFSCLICENT.CallMethod(1, "Chain.GetBlockByNumber", &req, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetBlockHashes sets the hash set of the specified scope chain blockchain
func (chain *ApiChain) GetBlockHashes(startHeight, endHeight string) (*respchain.Hashes, error) {

	if err := common.Str2Int64(startHeight); err != nil {
		return nil, err
	}
	if err := common.Str2Int64(endHeight); err != nil {
		return nil, err
	}
	req := &reqchain.GetBlockHashesArgs{
		Number: startHeight,
		Count:  endHeight,
	}

	result := new(respchain.Hashes)
	if err := apis.GVA_XFSCLICENT.CallMethod(1, "Chain.GetBlockHashes", &req, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetHead query the header of the highest block
func (chain *ApiChain) GetHead() (*respchain.BlockHeaderResp, error) {
	result := new(respchain.BlockHeaderResp)
	if err := apis.GVA_XFSCLICENT.CallMethod(1, "Chain.Head", nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetBlockHeaderByNumber specify the height of the blockchain and query blockheader information
func (chain *ApiChain) GetBlockHeaderByNumber(number string) (*respchain.BlockHeaderResp, error) {
	if err := common.Str2Int64(number); err != nil {
		return nil, err
	}
	req := &reqchain.GetBlockHeaderByNumberArgs{
		Number: number,
	}
	result := new(respchain.BlockHeaderResp)
	if err := apis.GVA_XFSCLICENT.CallMethod(1, "Chain.GetBlockHeaderByNumber", &req, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetBlockHeaderByHash specify block hash to query blockheader information
func (chain *ApiChain) GetBlockHeaderByHash(hash string) (*respchain.BlockHeaderResp, error) {

	req := &reqchain.GetBlockHeaderByHashArgs{
		Hash: hash,
	}
	result := new(respchain.BlockHeaderResp)
	if err := apis.GVA_XFSCLICENT.CallMethod(1, "Chain.GetBlockHeaderByHash", &req, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetBlockByHash specify block hash to query block information
func (chain *ApiChain) GetBlockByHash(hash string) (*respchain.BlockResp, error) {

	req := &reqchain.GetBlockByHashArgs{
		Hash: hash,
	}
	result := new(respchain.BlockResp)
	if err := apis.GVA_XFSCLICENT.CallMethod(1, "Chain.GetBlockByHash", &req, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetTxsByBlockNum specify the block height to obtain all transaction information of the block
func (chain *ApiChain) GetTxsByBlockNum(number string) (*resptransfer.TransactionsResp, error) {
	if err := common.Str2Int64(number); err != nil {
		return nil, err
	}
	req := &reqchain.GetTxsByBlockNumArgs{
		Number: number,
	}
	result := new(resptransfer.TransactionsResp)
	if err := apis.GVA_XFSCLICENT.CallMethod(1, "Chain.GetTxsByBlockNum", &req, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetTxsByBlockHash the specified block hash obtains all transaction information of the height
func (chain *ApiChain) GetTxsByBlockHash(hash string) (*resptransfer.TransactionsResp, error) {

	req := &reqchain.GetTxbyBlockHashArgs{
		Hash: hash,
	}
	result := new(resptransfer.TransactionsResp)
	if err := apis.GVA_XFSCLICENT.CallMethod(1, "Chain.GetTxsByBlockHash", &req, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetReceiptByHash specify transaction txhash to obtain transaction receipt information
func (chain *ApiChain) GetReceiptByHash(txhash string) (*resptransfer.ReceiptResp, error) {

	req := &reqchain.GetReceiptByHashArgs{
		Hash: txhash,
	}
	result := new(resptransfer.ReceiptResp)
	if err := apis.GVA_XFSCLICENT.CallMethod(1, "Chain.GetReceiptByHash", &req, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetTransaction specify transaction txhash to obtain transaction information
func (chain *ApiChain) GetTransaction(txhash string) (*resptransfer.TransactionResp, error) {

	req := &reqchain.GetTransactionArgs{
		Hash: txhash,
	}
	result := new(resptransfer.TransactionResp)
	if err := apis.GVA_XFSCLICENT.CallMethod(1, "Chain.GetTransaction", &req, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (chain *ApiChain) GetChainTransfer(txhash string) bool {
	req := &reqchain.GetTransactionArgs{
		Hash: txhash,
	}
	result := new(resptransfer.TransactionResp)
	if err := apis.GVA_XFSCLICENT.CallMethod(1, "Chain.GetTransaction", &req, &result); err != nil {
		return false
	}
	if result != nil {
		return true
	}
	return false
}

// GetBlockTxCountByHash get the number of all transactions in the specified block hash
func (chain *ApiChain) GetBlockTxCountByHash(hash string) (*int, error) {
	req := &reqchain.GetBlockTxCountByHashArgs{
		Hash: hash,
	}
	var result *int
	if err := apis.GVA_XFSCLICENT.CallMethod(1, "Chain.GetBlockTxCountByHash", &req, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetBlockTxCountByNum get the number of all transactions in the specified block height
func (chain *ApiChain) GetBlockTxCountByNum(number string) (*int, error) {
	if err := common.Str2Int64(number); err != nil {
		return nil, err
	}
	req := &reqchain.GetBlockTxCountByNumArgs{
		Number: number,
	}
	var result *int
	if err := apis.GVA_XFSCLICENT.CallMethod(1, "Chain.GetBlockTxCountByHash", &req, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetBlockTxByHashAndIndex query the transaction information with the specified block hash and the index of the specified transaction set
func (chain *ApiChain) GetBlockTxByHashAndIndex(hash string, index int) (*resptransfer.TransactionResp, error) {

	req := &reqchain.GetBlockTxByHashAndIndexArgs{
		Hash:  hash,
		Index: index,
	}
	result := new(resptransfer.TransactionResp)
	if err := apis.GVA_XFSCLICENT.CallMethod(1, "Chain.GetBlockTxCountByHash", &req, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetBlockTxByNumAndIndex query the transaction information by specifying the block height and the index of the specified transaction set
func (chain *ApiChain) GetBlockTxByNumAndIndex(number string, index int) (*resptransfer.TransactionResp, error) {
	if err := common.Str2Int64(number); err != nil {
		return nil, err
	}
	req := &reqchain.GetBlockTxByNumAndIndexArgs{
		Number: number,
		Index:  index,
	}

	result := new(resptransfer.TransactionResp)
	if err := apis.GVA_XFSCLICENT.CallMethod(1, "Chain.GetBlockTxCountByHash", &req, &result); err != nil {
		return nil, err
	}
	return result, nil
}
