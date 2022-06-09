package apicontract

import (
	"fmt"
	"math/big"
	"strconv"

	"github.com/younamebert/xfssdk/common"
	"github.com/younamebert/xfssdk/core/abi"
	"github.com/younamebert/xfssdk/core/apis"
	servetxpool "github.com/younamebert/xfssdk/servce/txpool"
)

type TokenLink interface {
	Create(TokenArgs) (string, error)
	// Name(args xfsgo.TokenCallRequest) (string, error)
	// Symbol(args xfsgo.TokenCallRequest) (string, error)
	// GetTotalSupply(args xfsgo.TokenCallRequest) (string, error)
	// BalanceOf(args xfsgo.TokenCallRequest) (string, error)
	// Mint(args xfsgo.TokenCallRequest) (string, error)
}

type ApiToken struct{}

func (token *ApiToken) Create(args TokenArgs) (string, error) {

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

func (token *ApiToken) DeployToken(args DeployTokenArgs) (string, error) {
	tokenTransfer := new(servetxpool.StringRawTransaction)
	tokenTransfer.GasLimit = common.TxGas.String()
	tokenTransfer.GasPrice = common.DefaultGasPrice().String()
	tokenTransfer.Data = args.Code

	address, err := common.StrKey2Address(args.Addresskey)
	if err != nil {
		return "", err
	}

	req := &GetAddrNonceByHashArgs{
		Address: address.B58String(),
	}
	var nonce *int64
	if err := apis.GVA_XFSCLICENT.CallMethod(1, "TxPool.GetAddrTxNonce", &req, &nonce); err != nil {
		return "", err
	}

	tokenTransfer.Nonce = strconv.FormatInt(*nonce, 10)

	if err := tokenTransfer.SignWithPrivateKey(args.Addresskey); err != nil {
		return "", err
	}

	// var txhash *string
	// if err := apis.GVA_XFSCLICENT.CallMethod(1, "TxPool.SendRawTransaction", &req, &txhash); err != nil {
	// 	return "", err
	// }
	return tokenTransfer.Transfer2Raw()
}
