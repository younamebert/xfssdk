package stdtoken

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

type StdTokenLocal struct{}

func (stdtokenlocad *StdTokenLocal) Create(args reqcontract.TokenArgs) (string, error) {

	decimal, ok := new(big.Int).SetString(args.Decimals, 10)
	if !ok {
		return "", fmt.Errorf("invalid decimals type string to big.Int error val:%v", args.Decimals)
	}

	totalSupply, ok := new(big.Int).SetString(args.TotalSupply, 10)
	if !ok {
		return "", fmt.Errorf("invalid totalSupply type string to big.Int error val:%v", args.TotalSupply)
	}
	code, err := apis.GVA_ABI_STDTOKEN.Create(abi.CTypeString(args.Name), abi.CTypeString(args.Symbol), abi.NewUint8(uint8(decimal.Uint64())), abi.NewUint256(totalSupply))
	if err != nil {
		return "", fmt.Errorf("an exception occurred of contract argument err:%v", err)
	}

	return code, err
}

func (stdtokenlocad *StdTokenLocal) DeployToken(args reqcontract.DeployTokenArgs) (*stdtoken, string, error) {
	//创建合约交易
	tokenTransfer := new(reqtransfer.StringRawTransaction)
	tokenTransfer.Data = args.Code

	address, err := libs.StrKey2Address(args.Addresskey)
	if err != nil {
		return nil, "", err
	}

	tokenTransfer.Value = "1"
	//签名交易
	stdtokentransfer, err := transfer.EnCodeRawTransaction(args.Addresskey, tokenTransfer)
	if err != nil {
		return nil, "", err
	}

	transfer2Raw, err := stdtokentransfer.Transfer2Raw()
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
	result := &stdtoken{
		ContractAddress:      caddr.B58String(), //合约地址
		CreatorAddressPrikey: args.Addresskey,   //创建人私钥
	}
	return result, txhash, nil
}

type StdTokenCall interface {
	Address() common.Address
	Name() (string, error)
	Symbol() (string, error)
	GetDecimals() (*big.Int, error)
	GetTotalSupply() (*big.Int, error)
	BalanceOf(account_address string) (*big.Int, error)
	Allowance(args reqcontract.StdTokenAllowanceArgs) (*big.Int, error)
	Mint(args reqcontract.StdTokenMintArgs) (string, error)
	Burn(args reqcontract.StdTokenBurnArgs) (string, error)
	Approve(args reqcontract.StdTokenApproveArgs) (string, error)
	TransferFrom(args reqcontract.StdTokenTransferFromArgs) (string, error)
}

type stdtoken struct {
	ContractAddress      string //合约地址
	CreatorAddressPrikey string //创建人私钥
}

func (stdtoken *stdtoken) Address() common.Address {
	address, err := libs.StrKey2Address(stdtoken.CreatorAddressPrikey)
	if err != nil {
		fmt.Printf("StrKey2Address :%v\n", err)
		os.Exit(1)
	}
	return address
}

//获取合约名称合约
func (stdtoken *stdtoken) Name() (string, error) {

	packed, err := apis.GVA_ABI_STDTOKEN.Name()
	if err != nil {
		return "", fmt.Errorf("no connection established in service err:%v", err)
	}
	req := contract.VMCallData{
		To:   stdtoken.ContractAddress,
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

func (stdtoken *stdtoken) Symbol() (string, error) {

	packed, err := apis.GVA_ABI_STDTOKEN.Symbol()
	if err != nil {
		return "", fmt.Errorf("no connection established in service err:%v", err)
	}
	req := contract.VMCallData{
		To:   stdtoken.ContractAddress,
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

func (stdtoken *stdtoken) GetDecimals() (*big.Int, error) {

	packed, err := apis.GVA_ABI_STDTOKEN.GetDecimals()
	if err != nil {
		return nil, fmt.Errorf("no connection established in service err:%v", err)
	}
	req := contract.VMCallData{
		To:   stdtoken.ContractAddress,
		Data: packed,
	}

	var result string
	if err := apis.GVA_XFSCLICENT.CallMethod(1, "VM.Call", &req, &result); err != nil {
		return nil, err
	}
	byteResult, err := common.HexToBytes(result)
	if err != nil {
		return nil, err
	}
	byteResult, err = common.HexToBytes(result)
	if err != nil {
		return big.NewInt(0), err
	}
	bigResult := big.NewInt(0).SetBytes(byteResult)
	return bigResult, nil
}

func (stdtoken *stdtoken) GetTotalSupply() (*big.Int, error) {

	packed, err := apis.GVA_ABI_STDTOKEN.GetTotalSupply()
	if err != nil {
		return nil, fmt.Errorf("no connection established in service err:%v", err)
	}
	req := contract.VMCallData{
		To:   stdtoken.ContractAddress,
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

func (stdtoken *stdtoken) BalanceOf(account_address string) (*big.Int, error) {

	accoutAddr := common.StrB58ToAddress(account_address)

	cTypeAddr := abi.NewAddress(accoutAddr)
	packed, err := apis.GVA_ABI_STDTOKEN.BalanceOf(cTypeAddr)
	if err != nil {
		return nil, fmt.Errorf("no connection established in service err:%v", err)
	}

	req := contract.VMCallData{
		To:   stdtoken.ContractAddress,
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

func (stdtoken *stdtoken) Allowance(args reqcontract.StdTokenAllowanceArgs) (*big.Int, error) {

	ownerAddr := common.StrB58ToAddress(args.OwnerAddress)
	cTypeOwnerAddr := abi.NewAddress(ownerAddr)

	addr := common.StrB58ToAddress(args.SpenderAddress)
	cTypeAddr := abi.NewAddress(addr)

	packed, err := apis.GVA_ABI_STDTOKEN.Allowance(cTypeOwnerAddr, cTypeAddr)
	if err != nil {
		return nil, fmt.Errorf("no connection established in service err:%v", err)
	}

	req := contract.VMCallData{
		To:   stdtoken.ContractAddress,
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

func (stdtoken *stdtoken) Mint(args reqcontract.StdTokenMintArgs) (string, error) {

	address, err := libs.StrKey2Address(stdtoken.CreatorAddressPrikey)
	if err != nil {
		return "", fmt.Errorf("invalid MintAddressPriKey to address err:%v", err)
	}

	cTypeAddr := abi.NewAddress(address)
	amount, ok := new(big.Int).SetString(args.Amount, 10)
	if !ok {
		return "", fmt.Errorf("invalid Mint error")
	}
	packed, err := apis.GVA_ABI_STDTOKEN.Mint(cTypeAddr, abi.NewUint256(amount))
	if err != nil {
		return "", fmt.Errorf("no connection established in service err:%v", err)
	}

	tokenTransfer := new(reqtransfer.StringRawTransaction)
	//初始化GAS和code
	tokenTransfer.To = stdtoken.ContractAddress
	tokenTransfer.Data = packed

	stdtokentransfer, err := transfer.EnCodeRawTransaction(stdtoken.CreatorAddressPrikey, tokenTransfer)
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

func (stdtoken *stdtoken) Burn(args reqcontract.StdTokenBurnArgs) (string, error) {

	addr := common.StrB58ToAddress(args.BurnAddress)
	cTypeAddr := abi.NewAddress(addr)
	amount, ok := new(big.Int).SetString(args.Amount, 10)
	if !ok {
		return "", fmt.Errorf("invalid Mint error")
	}
	packed, err := apis.GVA_ABI_STDTOKEN.Burn(cTypeAddr, abi.NewUint256(amount))
	if err != nil {
		return "", fmt.Errorf("no connection established in service err:%v", err)
	}

	tokenTransfer := new(reqtransfer.StringRawTransaction)
	//初始化GAS和code
	tokenTransfer.To = args.BurnAddress
	tokenTransfer.Data = packed
	stdtokentransfer, err := transfer.EnCodeRawTransaction(stdtoken.CreatorAddressPrikey, tokenTransfer)
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

func (stdtoken *stdtoken) Approve(args reqcontract.StdTokenApproveArgs) (string, error) {

	addr := common.StrB58ToAddress(args.ApproveSpenderAddress)
	cTypeAddr := abi.NewAddress(addr)
	amount, ok := new(big.Int).SetString(args.Amount, 10)
	if !ok {
		return "", fmt.Errorf("invalid Mint error")
	}
	packed, err := apis.GVA_ABI_STDTOKEN.Approve(cTypeAddr, abi.NewUint256(amount))
	if err != nil {
		return "", fmt.Errorf("no connection established in service err:%v", err)
	}

	tokenTransfer := new(reqtransfer.StringRawTransaction)
	//初始化GAS和code
	tokenTransfer.To = args.ApproveSpenderAddress
	tokenTransfer.Data = packed
	stdtokentransfer, err := transfer.EnCodeRawTransaction(args.ApproveFromAddressPriKey, tokenTransfer)
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

func (stdtoken *stdtoken) TransferFrom(args reqcontract.StdTokenTransferFromArgs) (string, error) {
	fromAddr := common.StrB58ToAddress(args.TransferFromAddress)
	toAddr := common.StrB58ToAddress(args.TransferToAddress)

	amount, ok := new(big.Int).SetString(args.TransferFromValue, 10)
	if !ok {
		return "", fmt.Errorf("invalid Mint error")
	}
	packed, err := apis.GVA_ABI_STDTOKEN.TransferFrom(abi.NewAddress(fromAddr), abi.NewAddress(toAddr), abi.NewUint256(amount))
	if err != nil {
		return "", fmt.Errorf("no connection established in service err:%v", err)
	}

	tokenTransfer := new(reqtransfer.StringRawTransaction)
	//初始化GAS和code
	tokenTransfer.To = args.TransferToAddress
	tokenTransfer.Data = packed
	stdtokentransfer, err := transfer.EnCodeRawTransaction(args.TransferSpenderAddressPriKey, tokenTransfer)
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
