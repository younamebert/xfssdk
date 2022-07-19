package bridge

import (
	"fmt"
	"math/big"
	"strconv"

	"github.com/younamebert/xfssdk/common"
	"github.com/younamebert/xfssdk/common/ahash"
	"github.com/younamebert/xfssdk/core/abi"
	"github.com/younamebert/xfssdk/core/apis"
	"github.com/younamebert/xfssdk/core/transfer"
	"github.com/younamebert/xfssdk/crypto"
	"github.com/younamebert/xfssdk/libs"
	reqcontract "github.com/younamebert/xfssdk/servce/contract/request"
	reqtransfer "github.com/younamebert/xfssdk/servce/transfer/request"
)

type BridgeLocal struct{}

func (bridgelocal *BridgeLocal) Create(args *reqcontract.BridgeArgs) (string, error) {
	code, err := apis.GVA_ABI_BRIDGETOKEN.Create(abi.CTypeString(args.Name), abi.CTypeString(args.Symbol), abi.NewAddress(common.StrB58ToAddress(args.ContractAddress)), abi.NewUint256(args.ChainId))
	if err != nil {
		return "", fmt.Errorf("an exception occurred of contract argument err:%v", err)
	}
	return code, err
}

func (bridgelocal *BridgeLocal) DeployDridge(args reqcontract.DeployBridgeArgs) (*Bridge, string, error) {
	//创建合约交易
	tokenTransfer := new(reqtransfer.StringRawTransaction)
	tokenTransfer.Data = args.Code

	address, err := libs.StrKey2Address(args.Addresskey)
	if err != nil {
		return nil, "", err
	}
	//签名交易
	stdtokentransfer, err := transfer.EnCodeRawTransaction(args.Addresskey, tokenTransfer)
	if err != nil {
		return nil, "", err
	}

	transfer2Raw, err := stdtokentransfer.Transfer2Raw()
	fmt.Println(transfer2Raw)
	if err != nil {
		return nil, "", err
	}
	// 发送一笔合约交易
	txhash, err := transfer.SendRawTransfer(transfer2Raw)
	if err != nil {
		return nil, "", err
	}

	//创建合约地址
	fromAddressHashBytes := ahash.SHA256(address[:])
	fromAddressHash := common.Bytes2Hash(fromAddressHashBytes)

	nonce, err := strconv.ParseUint(stdtokentransfer.Nonce, 10, 64)
	if err != nil {
		return nil, "", err
	}
	caddr := crypto.CreateAddress(fromAddressHash, nonce)
	//返回交易哈希和合约地址
	result := &Bridge{
		Bankaddress:          caddr.B58String(), //合约地址
		CreatorAddressPrikey: args.Addresskey,   //创建人私钥
	}
	return result, txhash, nil
}

type Bridge struct {
	Bankaddress          string //合约地址
	CreatorAddressPrikey string //创建人私钥
}

func (bridge *Bridge) TransferIn(args reqcontract.BridgeTransferInArgs) (string, error) {
	toAddr := common.StrB58ToAddress(args.TransferToAddress)
	amount, ok := new(big.Int).SetString(args.TransferAmount, 10)
	if !ok {
		return "", fmt.Errorf("invalid Mint error")
	}
	toChainId, ok := new(big.Int).SetString(args.TransferToChainId, 10)
	if !ok {
		return "", fmt.Errorf("invalid Mint error")
	}
	packed, err := apis.GVA_ABI_BRIDGETOKEN.TransferIn(abi.NewAddress(toAddr), abi.NewUint256(amount), abi.NewUint256(toChainId))
	if err != nil {
		return "", fmt.Errorf("no connection established in service err:%v", err)
	}
	tokenTransfer := new(reqtransfer.StringRawTransaction)
	tokenTransfer.To = bridge.Bankaddress
	tokenTransfer.Data = packed
	stdtokentransfer, err := transfer.EnCodeRawTransaction(args.TransferFromAddressPriKey, tokenTransfer)
	if err != nil {
		return "", fmt.Errorf("no connection established in service err:%v", err)
	}
	//交易数据转base64
	transfer2Raw, err := stdtokentransfer.Transfer2Raw()
	if err != nil {
		return "", err
	}
	return transfer.SendRawTransfer(transfer2Raw)
}
