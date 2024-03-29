package nfttoken

import (
	"fmt"
	"math/big"
	"os"
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

type NFTTokenLocal struct{}

func (nftokenlocad *NFTTokenLocal) NFTCreate(args reqcontract.NFTTokenCreateArgs) (string, error) {
	code, err := apis.GVA_ABI_NFTTOKEN.Create(abi.CTypeString(args.Name), abi.CTypeString(args.Symbol))
	if err != nil {
		return "", fmt.Errorf("an exception occurred of contract argument err:%v", err)
	}
	return code, nil
}

func (nftokenlocad *NFTTokenLocal) NFTDeployToken(args reqcontract.DeployNFTokenArgs) (NFTTokenCall, string, error) {
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
	if err != nil {
		return nil, "", err
	}

	bs, _ := common.MarshalIndent(transfer2Raw)
	fmt.Println(string(bs))
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
	result := &NFToken{
		ContractAddress:      caddr.B58String(), //合约地址
		CreatorAddressPrikey: args.Addresskey,   //创建人私钥
	}
	return result, txhash, nil
}

type NFTTokenCall interface {
	Name() (string, error)
	Symbol() (string, error)
	BalanceOf(args reqcontract.NFTBalanceOfArgs) (*big.Int, error)
	OwnerOf(args reqcontract.NFTOwnerOfArgs) (string, error)
	IsApproveForAll(args reqcontract.NFTISApproveForAllArgs) (string, error)
	GetApproved(args reqcontract.NFTGetApprovedArgs) (string, error)
	Mint(args reqcontract.NFTokenMintArgs) (string, error)
	Approve(args reqcontract.NFTApproveArgs) (string, error)
	SetApprovalForAll(args reqcontract.NFTSetApproveForAllArgs) (string, error)
	TransferFrom(args reqcontract.NFTTransferFromArgs) (string, error)
}

type NFToken struct {
	ContractAddress      string //合约地址
	CreatorAddressPrikey string //创建人私钥
}

func (nfttoken *NFToken) Address() common.Address {
	address, err := libs.StrKey2Address(nfttoken.CreatorAddressPrikey)
	if err != nil {
		fmt.Printf("StrKey2Address :%v\n", err)
		os.Exit(1)
	}
	return address
}

//获取合约名称合约
func (nfttoken *NFToken) Name() (string, error) {

	packed, err := apis.GVA_ABI_NFTTOKEN.Name()
	if err != nil {
		return "", fmt.Errorf("no connection established in service err:%v", err)
	}
	req := contract.VMCallData{
		To:   nfttoken.ContractAddress,
		Data: packed,
	}
	var result string
	if err := apis.GVA_XFSCLICENT.CallMethod(1, "VM.Call", &req, &result); err != nil {
		return "", err
	}
	byteResult, err := common.HexToBytes(result)
	if err != nil {
		return "", err
	}
	tokenname := string(byteResult)
	return tokenname, nil
}

func (nfttoken *NFToken) Symbol() (string, error) {

	packed, err := apis.GVA_ABI_NFTTOKEN.Symbol()
	if err != nil {
		return "", fmt.Errorf("no connection established in service err:%v", err)
	}
	req := contract.VMCallData{
		To:   nfttoken.ContractAddress,
		Data: packed,
	}
	var result string
	if err := apis.GVA_XFSCLICENT.CallMethod(1, "VM.Call", &req, &result); err != nil {
		return "", err
	}
	byteResult, err := common.HexToBytes(result)
	if err != nil {
		return "", err
	}
	tokenname := string(byteResult)
	return tokenname, nil
}

func (nfttoken *NFToken) BalanceOf(args reqcontract.NFTBalanceOfArgs) (*big.Int, error) {
	packed, err := apis.GVA_ABI_NFTTOKEN.BalanceOf(args.BalanceOfAddress)
	if err != nil {
		return nil, fmt.Errorf("no connection established in service err:%v", err)
	}
	req := contract.VMCallData{
		To:   nfttoken.ContractAddress,
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

func (nfttoken *NFToken) OwnerOf(args reqcontract.NFTOwnerOfArgs) (string, error) {

	tokenid, ok := new(big.Int).SetString(args.TokenId, 10)
	if !ok {
		return "", fmt.Errorf("invalid TokenId on the %v", args.TokenId)
	}

	packed, err := apis.GVA_ABI_NFTTOKEN.OwnerOf(abi.NewUint256(tokenid))
	if err != nil {
		return "", fmt.Errorf("no connection established in service err:%v", err)
	}
	req := contract.VMCallData{
		To:   nfttoken.ContractAddress,
		Data: packed,
	}
	var result string
	if err := apis.GVA_XFSCLICENT.CallMethod(1, "VM.Call", &req, &result); err != nil {
		return "", err
	}

	byteData, err := common.HexToBytes(result)
	if err != nil {
		return "", err
	}

	addr := common.Bytes2Address(byteData)

	return addr.B58String(), nil
}

func (nfttoken *NFToken) IsApproveForAll(args reqcontract.NFTISApproveForAllArgs) (string, error) {

	ownerAddr := common.StrB58ToAddress(args.OwnerAddress)
	spenderAddr := common.StrB58ToAddress(args.SpenderAddress)

	packed, err := apis.GVA_ABI_NFTTOKEN.IsApproveForAll(abi.NewAddress(ownerAddr), abi.NewAddress(spenderAddr))
	if err != nil {
		return "", fmt.Errorf("no connection established in service err:%v", err)
	}

	req := contract.VMCallData{
		To:   nfttoken.ContractAddress,
		Data: packed,
	}

	var result string
	if err := apis.GVA_XFSCLICENT.CallMethod(1, "VM.Call", &req, &result); err != nil {
		return "", err
	}

	byteResult, err := common.HexToBytes(result)
	if err != nil {
		return "", err
	}
	strData := string(byteResult)

	return strData, nil
}

func (nfttoken *NFToken) GetApproved(args reqcontract.NFTGetApprovedArgs) (string, error) {

	tokenid, ok := new(big.Int).SetString(args.TokenId, 10)
	if !ok {
		return "", fmt.Errorf("invalid tokenId on the error")
	}

	packed, err := apis.GVA_ABI_NFTTOKEN.GetApproved(abi.NewUint256(tokenid))
	if err != nil {
		return "", fmt.Errorf("no connection established in service err:%v", err)
	}

	req := contract.VMCallData{
		To:   nfttoken.ContractAddress,
		Data: packed,
	}

	var result string
	if err := apis.GVA_XFSCLICENT.CallMethod(1, "VM.Call", &req, &result); err != nil {
		return "", err
	}

	byteData, err := common.HexToBytes(result)
	if err != nil {
		return "", err
	}

	addr := common.Bytes2Address(byteData)

	return addr.B58String(), nil
}

func (nfttoken *NFToken) Mint(args reqcontract.NFTokenMintArgs) (string, error) {

	address, err := libs.StrKey2Address(nfttoken.CreatorAddressPrikey)
	if err != nil {
		return "", fmt.Errorf("invalid MintAddressPriKey to address err:%v", err)
	}

	packed, err := apis.GVA_ABI_NFTTOKEN.Mint(abi.NewAddress(address), abi.CTypeString(args.TokenId))
	if err != nil {
		return "", fmt.Errorf("no connection established in service err:%v", err)
	}
	tokenTransfer := new(reqtransfer.StringRawTransaction)
	//初始化GAS和code
	tokenTransfer.To = nfttoken.ContractAddress
	tokenTransfer.Data = packed
	// fmt.Printf()
	stdtokentransfer, err := transfer.EnCodeRawTransaction(nfttoken.CreatorAddressPrikey, tokenTransfer)
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

func (nfttoken *NFToken) Approve(args reqcontract.NFTApproveArgs) (string, error) {

	addressTo := common.StrB58ToAddress(args.ApproveToAddress)

	openid, ok := new(big.Int).SetString(args.Openid, 10)
	if !ok {
		return "", fmt.Errorf("invalid tokenId on the error")
	}

	packed, err := apis.GVA_ABI_NFTTOKEN.Approve(abi.NewAddress(addressTo), abi.NewUint256(openid))
	if err != nil {
		return "", fmt.Errorf("no connection established in service err:%v", err)
	}

	tokenTransfer := new(reqtransfer.StringRawTransaction)
	//初始化GAS和code
	tokenTransfer.To = args.ApproveToAddress
	tokenTransfer.Data = packed

	stdtokentransfer, err := transfer.EnCodeRawTransaction(args.ApproveFromAddressPriKey, tokenTransfer)
	if err != nil {
		return "", fmt.Errorf("invalid mint err:%v", err)
	}

	//交易数据转base64
	transfer2Raw, err := stdtokentransfer.Transfer2Raw()
	if err != nil {
		return "", err
	}
	return transfer.SendRawTransfer(transfer2Raw)
}

func (nfttoken *NFToken) SetApprovalForAll(args reqcontract.NFTSetApproveForAllArgs) (string, error) {

	toAddr := common.StrB58ToAddress(args.ApproveallToAddress)

	var cTypeStatus abi.CTypeBool
	if args.AllApproved {
		cTypeStatus = abi.CTypeBool{1}
	} else {
		cTypeStatus = abi.CTypeBool{0}
	}

	packed, err := apis.GVA_ABI_NFTTOKEN.SetApprovalForAll(abi.NewAddress(toAddr), cTypeStatus)
	if err != nil {
		return "", fmt.Errorf("no connection established in service err:%v", err)
	}

	tokenTransfer := new(reqtransfer.StringRawTransaction)
	//初始化GAS和code
	tokenTransfer.To = args.ApproveallToAddress
	tokenTransfer.Data = packed

	stdtokentransfer, err := transfer.EnCodeRawTransaction(args.ApproveallFromAddressPriKey, tokenTransfer)
	if err != nil {
		return "", fmt.Errorf("invalid mint err:%v", err)
	}

	//交易数据转base64
	transfer2Raw, err := stdtokentransfer.Transfer2Raw()
	if err != nil {
		return "", err
	}
	return transfer.SendRawTransfer(transfer2Raw)
}

func (nfttoken *NFToken) TransferFrom(args reqcontract.NFTTransferFromArgs) (string, error) {

	fromAddr := common.StrB58ToAddress(args.TransferFromAddress)
	toAddr := common.StrB58ToAddress(args.TransferToAddress)

	tokenid, ok := new(big.Int).SetString(args.TransferFromTokenId, 10)
	if !ok {
		return "", fmt.Errorf("invalid Mint error")
	}
	packed, err := apis.GVA_ABI_STDTOKEN.TransferFrom(abi.NewAddress(fromAddr), abi.NewAddress(toAddr), abi.NewUint256(tokenid))
	if err != nil {
		return "", fmt.Errorf("no connection established in service err:%v", err)
	}

	tokenTransfer := new(reqtransfer.StringRawTransaction)
	//初始化GAS和code
	tokenTransfer.To = nfttoken.ContractAddress
	tokenTransfer.Data = packed
	stdtokentransfer, err := transfer.EnCodeRawTransaction(args.TransferOperatorAddressPriKey, tokenTransfer)
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
