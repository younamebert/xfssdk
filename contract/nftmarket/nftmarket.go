package nftmarket

import (
	"errors"
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
	"github.com/younamebert/xfssdk/servce/contract"
	reqcontract "github.com/younamebert/xfssdk/servce/contract/request"
	reqtransfer "github.com/younamebert/xfssdk/servce/transfer/request"
)

type NftMarketLocal struct{}

func (nftmarketlocal *NftMarketLocal) Create(args *reqcontract.NftMarketArgs) (string, error) {
	code, err := apis.GVA_ABI_NFTMARKET.Create()
	if err != nil {
		return "", fmt.Errorf("an exception occurred of contract argument err:%v", err)
	}
	return code, err
}

func (nftmarketlocal *NftMarketLocal) Deploy(args reqcontract.DeployNftMarketArgs) (*NftMarket, string, error) {
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
	result := &NftMarket{
		ContractAddress:      caddr.B58String(), //合约地址
		CreatorAddressPrikey: args.Addresskey,   //创建人私钥
	}
	return result, txhash, nil
}

type NftMarket struct {
	ContractAddress      string //合约地址
	CreatorAddressPrikey string //创建人私钥
}

func (nftmarket *NftMarket) Mint(args *reqcontract.NFTMarketMintArgs) (string, error) {

	address := common.StrB58ToAddress(args.Address)
	amount := abi.NewUint256(args.Amount)
	tokenid := abi.CTypeString(args.TokenUrl)
	packed, err := apis.GVA_ABI_NFTMARKET.Mint(abi.NewAddress(address), amount, tokenid)
	if err != nil {
		return "", fmt.Errorf("no connection established in service err:%v", err)
	}
	tokenTransfer := new(reqtransfer.StringRawTransaction)
	//初始化GAS和code
	tokenTransfer.To = nftmarket.ContractAddress
	tokenTransfer.Data = packed
	// fmt.Printf()
	stdtokentransfer, err := transfer.EnCodeRawTransaction(nftmarket.CreatorAddressPrikey, tokenTransfer)
	if err != nil {
		return "", fmt.Errorf("invalid mint err:%v", err)
	}

	//交易数据转base64
	transfer2Raw, err := stdtokentransfer.Transfer2Raw()
	fmt.Println(transfer2Raw)
	if err != nil {
		return "", err
	}
	return transfer.SendRawTransfer(transfer2Raw)
}

func (nftmarket *NftMarket) MintBatch(args *reqcontract.NFTMarketMintBatchArgs) (string, error) {
	address := common.StrB58ToAddress(args.Address)

	if len(args.Amounts) != len(args.TokenUrls) {
		return "", errors.New("Amounts.Len != tokenurls.Len")
	}
	amounts := make([]abi.CTypeUint256, 0)
	tokenUrls := make([]abi.CTypeString, 0)
	for i := 0; i < len(args.TokenUrls); i++ {
		amount := abi.NewUint256(args.Amounts[i])
		amounts = append(amounts, amount)

		tokenUrl := abi.CTypeString(args.TokenUrls[i])
		tokenUrls = append(tokenUrls, tokenUrl)
	}
	packed, err := apis.GVA_ABI_NFTMARKET.Mint(abi.NewAddress(address), amounts, tokenUrls)
	if err != nil {
		return "", fmt.Errorf("no connection established in service err:%v", err)
	}
	tokenTransfer := new(reqtransfer.StringRawTransaction)
	//初始化GAS和code
	tokenTransfer.To = nftmarket.ContractAddress
	tokenTransfer.Data = packed
	// fmt.Printf()
	stdtokentransfer, err := transfer.EnCodeRawTransaction(nftmarket.CreatorAddressPrikey, tokenTransfer)
	if err != nil {
		return "", fmt.Errorf("invalid mint err:%v", err)
	}

	//交易数据转base64
	transfer2Raw, err := stdtokentransfer.Transfer2Raw()
	fmt.Println(transfer2Raw)
	if err != nil {
		return "", err
	}
	return transfer.SendRawTransfer(transfer2Raw)
}

func (nftmarket *NftMarket) BalanceOf(account_address string, tokenid *big.Int) (*big.Int, error) {
	accoutAddr := common.StrB58ToAddress(account_address)
	cTypeAddr := abi.NewAddress(accoutAddr)
	id := abi.NewUint256(tokenid)
	packed, err := apis.GVA_ABI_NFTMARKET.BalanceOf(cTypeAddr, id)
	if err != nil {
		return nil, fmt.Errorf("no connection established in service err:%v", err)
	}

	req := contract.VMCallData{
		To:   nftmarket.ContractAddress,
		Data: packed,
	}
	var result string
	if err := apis.GVA_XFSCLICENT.CallMethod(1, "VM.Call", &req, &result); err != nil {
		return nil, err
	}
	byteResult, err := common.HexToBytes(result)
	if err != nil {
		return big.NewInt(0), err
	}
	bigResult := big.NewInt(0).SetBytes(byteResult)
	return bigResult, nil
}
