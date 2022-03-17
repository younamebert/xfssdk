package apichain

import (
	"xfssdk/common"
	"xfssdk/core/apis"
)

type ChainLink interface {
	GetBlockByNumber(number string) (*apis.BlockResp, error)
	GetBlockHashes(startHeight, endHeight string) (*apis.Hashes, error)
	GetHead() (*apis.BlockHeaderResp, error)
	GetBlockHeaderByNumber(number string) (*apis.BlockHeaderResp, error)
	GetBlockHeaderByHash(hash string) (*apis.BlockHeaderResp, error)
	GetBlockByHash(hash string) (*apis.BlockResp, error)
	GetTxsByBlockNum(number string) (*apis.TransactionsResp, error)
	GetTxsByBlockHash(hash string) (*apis.TransactionsResp, error)
	GetReceiptByHash(hash string) (*apis.ReceiptResp, error)
	GetTransaction(txhash string) (*apis.TransactionResp, error)
	GetBlockTxCountByHash(hash string) (*int, error)
	GetBlockTxCountByNum(number string) (*int, error)
	GetBlockTxByHashAndIndex(hash string, index int) (*apis.TransactionResp, error)
	GetBlockTxByNumAndIndex(number string, index int) (*apis.TransactionResp, error)
}

type ApiChain struct{}

// func NewApiChain() *ApiChain {
// 	return ne
// }
//GetBlockByNumber 根据区块高度获取区块信息
func (chain *ApiChain) GetBlockByNumber(number string) (*apis.BlockResp, error) {

	if err := common.Str2Int64(number); err != nil {
		return nil, err
	}
	req := &GetBlockByNumArgs{
		Number: number,
	}
	result := new(apis.BlockResp)
	if err := apis.GVA_XFSCLICENT.CallMethod(1, "Chain.GetBlockByNumber", &req, &result); err != nil {
		return nil, err
	}
	return result, nil
}

//GetBlockHashes 获取指定范围链区块的hash集合
func (chain *ApiChain) GetBlockHashes(startHeight, endHeight string) (*apis.Hashes, error) {

	if err := common.Str2Int64(startHeight); err != nil {
		return nil, err
	}
	if err := common.Str2Int64(endHeight); err != nil {
		return nil, err
	}
	req := &GetBlockHashesArgs{
		Number: startHeight,
		Count:  endHeight,
	}

	result := new(apis.Hashes)
	if err := apis.GVA_XFSCLICENT.CallMethod(1, "Chain.GetBlockHashes", &req, &result); err != nil {
		return nil, err
	}
	return result, nil
}

//GetHead 获取链上最新的区块
func (chain *ApiChain) GetHead() (*apis.BlockHeaderResp, error) {
	result := new(apis.BlockHeaderResp)
	if err := apis.GVA_XFSCLICENT.CallMethod(1, "Chain.Head", nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (chain *ApiChain) GetBlockHeaderByNumber(number string) (*apis.BlockHeaderResp, error) {
	if err := common.Str2Int64(number); err != nil {
		return nil, err
	}
	req := &GetBlockHeaderByNumberArgs{
		Number: number,
	}
	result := new(apis.BlockHeaderResp)
	if err := apis.GVA_XFSCLICENT.CallMethod(1, "Chain.GetBlockHeaderByNumber", &req, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (chain *ApiChain) GetBlockHeaderByHash(hash string) (*apis.BlockHeaderResp, error) {

	req := &GetBlockHeaderByHashArgs{
		Hash: hash,
	}
	result := new(apis.BlockHeaderResp)
	if err := apis.GVA_XFSCLICENT.CallMethod(1, "Chain.GetBlockHeaderByHash", &req, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (chain *ApiChain) GetBlockByHash(hash string) (*apis.BlockResp, error) {

	req := &GetBlockByHashArgs{
		Hash: hash,
	}
	result := new(apis.BlockResp)
	if err := apis.GVA_XFSCLICENT.CallMethod(1, "Chain.GetBlockByHash", &req, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (chain *ApiChain) GetTxsByBlockNum(number string) (*apis.TransactionsResp, error) {
	if err := common.Str2Int64(number); err != nil {
		return nil, err
	}
	req := &GetTxsByBlockNumArgs{
		Number: number,
	}
	result := new(apis.TransactionsResp)
	if err := apis.GVA_XFSCLICENT.CallMethod(1, "Chain.GetTxsByBlockNum", &req, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (chain *ApiChain) GetTxsByBlockHash(hash string) (*apis.TransactionsResp, error) {

	req := &GetTxbyBlockHashArgs{
		Hash: hash,
	}
	result := new(apis.TransactionsResp)
	if err := apis.GVA_XFSCLICENT.CallMethod(1, "Chain.GetTxsByBlockHash", &req, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (chain *ApiChain) GetReceiptByHash(hash string) (*apis.ReceiptResp, error) {

	req := &GetReceiptByHashArgs{
		Hash: hash,
	}
	result := new(apis.ReceiptResp)
	if err := apis.GVA_XFSCLICENT.CallMethod(1, "Chain.GetReceiptByHash", &req, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (chain *ApiChain) GetTransaction(txhash string) (*apis.TransactionResp, error) {

	req := &GetTransactionArgs{
		Hash: txhash,
	}
	result := new(apis.TransactionResp)
	if err := apis.GVA_XFSCLICENT.CallMethod(1, "Chain.GetTransaction", &req, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// Syncing returns false in case the node is currently not syncing with the network. It can be up to date or has not
// yet received the latest block headers from its pears. In case it is synchronizing:
// - startingBlock: block number this node started to synchronise from
// - currentBlock:  block number this node is currently importing
// - highestBlock:  block number of the highest block header this node has received from peers
// func (chain *ApiChain) GetSyncStatus() (apis.ChainStatusResp, error) {
// 	result := new(apis.ChainStatusResp)
// 	if err := apis.GVA_XFSCLICENT.CallMethod(1, "Chain.GetSyncStatus", nil, &result); err != nil {
// 		return nil, err
// 	}
// 	return result, nil
// }

// func (chain *ApiChain) GetSyncStatus() (apis.ChainStatusResp, error) {
// 	result := new(apis.ChainStatusResp)
// 	if err := apis.GVA_XFSCLICENT.CallMethod(1, "Chain.GetSyncStatus", nil, &result); err != nil {
// 		return nil, err
// 	}
// 	return result, nil
// }

func (chain *ApiChain) GetBlockTxCountByHash(hash string) (*int, error) {
	req := &GetBlockTxCountByHashArgs{
		Hash: hash,
	}
	var result *int
	if err := apis.GVA_XFSCLICENT.CallMethod(1, "Chain.GetBlockTxCountByHash", &req, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (chain *ApiChain) GetBlockTxCountByNum(number string) (*int, error) {
	if err := common.Str2Int64(number); err != nil {
		return nil, err
	}
	req := &GetBlockTxCountByNumArgs{
		Number: number,
	}
	var result *int
	if err := apis.GVA_XFSCLICENT.CallMethod(1, "Chain.GetBlockTxCountByHash", &req, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (chain *ApiChain) GetBlockTxByHashAndIndex(hash string, index int) (*apis.TransactionResp, error) {

	req := &GetBlockTxByHashAndIndexArgs{
		Hash:  hash,
		Index: index,
	}
	result := new(apis.TransactionResp)
	if err := apis.GVA_XFSCLICENT.CallMethod(1, "Chain.GetBlockTxCountByHash", &req, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (chain *ApiChain) GetBlockTxByNumAndIndex(number string, index int) (*apis.TransactionResp, error) {
	if err := common.Str2Int64(number); err != nil {
		return nil, err
	}
	req := &GetBlockTxByNumAndIndexArgs{
		Number: number,
		Index:  index,
	}

	result := new(apis.TransactionResp)
	if err := apis.GVA_XFSCLICENT.CallMethod(1, "Chain.GetBlockTxCountByHash", &req, &result); err != nil {
		return nil, err
	}
	return result, nil
}
