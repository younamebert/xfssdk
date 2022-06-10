package contract

import (
	"fmt"
	"math/big"
	"strconv"

	"github.com/younamebert/xfssdk/common"
	"github.com/younamebert/xfssdk/common/ahash"
	"github.com/younamebert/xfssdk/contract/stdtoken"
	"github.com/younamebert/xfssdk/core/abi"
	"github.com/younamebert/xfssdk/core/apis"
	"github.com/younamebert/xfssdk/core/transfer"
	"github.com/younamebert/xfssdk/crypto"
	"github.com/younamebert/xfssdk/libs"
	reqcontract "github.com/younamebert/xfssdk/servce/contract/request"
	respcontract "github.com/younamebert/xfssdk/servce/contract/response"
	reqtransfer "github.com/younamebert/xfssdk/servce/transfer/request"
)

type ContractEngine struct {
	Stdtoken stdtoken.StdToken
}

func Create(args reqcontract.TokenArgs) (string, error) {

	decimal, ok := new(big.Int).SetString(args.Decimals, 10)
	if !ok {
		return "", fmt.Errorf("invalid decimals type string to big.Int error val:%v", args.Decimals)
	}

	totalSupply, ok := new(big.Int).SetString(args.TotalSupply, 10)
	if !ok {
		return "", fmt.Errorf("invalid totalSupply type string to big.Int error val:%v", args.TotalSupply)
	}

	code, err := apis.GVA_ABI.Create(abi.CTypeString(args.Name), abi.CTypeString(args.Symbol), abi.NewUint8(uint8(decimal.Uint64())), abi.NewUint256(totalSupply))
	if err != nil {
		return "", fmt.Errorf("an exception occurred of contract argument")
	}

	return code, err
}

func DeployToken(args reqcontract.DeployTokenArgs) (*respcontract.DeployTokenResp, error) {
	//创建合约交易
	tokenTransfer := new(reqtransfer.StringRawTransaction)
	tokenTransfer.Data = args.Code

	address, err := libs.StrKey2Address(args.Addresskey)
	if err != nil {
		return nil, err
	}
	//签名交易
	stdtokentransfer, err := transfer.EnCodeRawTransaction(args.Addresskey, tokenTransfer)
	if err != nil {
		return nil, err
	}

	transfer2Raw, err := stdtokentransfer.Transfer2Raw()
	if err != nil {
		return nil, err
	}
	// 发送一笔合约交易
	txhash, err := transfer.SendRawTransfer(transfer2Raw)
	if err != nil {
		return nil, err
	}

	//创建合约地址
	fromAddressHashBytes := ahash.SHA256(address[:])
	fromAddressHash := common.Bytes2Hash(fromAddressHashBytes)

	nonce, err := strconv.ParseUint(stdtokentransfer.Nonce, 10, 64)
	if err != nil {
		return nil, err
	}
	caddr := crypto.CreateAddress(fromAddressHash, nonce)

	//返回交易哈希和合约地址
	result := &respcontract.DeployTokenResp{
		TransferHash:    txhash,
		ContractAddress: caddr.B58String(),
	}
	return result, nil
}
