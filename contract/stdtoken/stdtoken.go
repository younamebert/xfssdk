package stdtoken

import (
	"fmt"
	"math/big"
	"os"

	"github.com/younamebert/xfssdk/common"
	"github.com/younamebert/xfssdk/core/abi"
	"github.com/younamebert/xfssdk/core/apis"
	"github.com/younamebert/xfssdk/core/transfer"
	"github.com/younamebert/xfssdk/libs"
	"github.com/younamebert/xfssdk/servce/contract"
	reqcontract "github.com/younamebert/xfssdk/servce/contract/request"
	reqtransfer "github.com/younamebert/xfssdk/servce/transfer/request"
)

type StdToken interface {
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

type Stdtoken struct {
	ContractAddress      string //合约地址
	CreatorAddressPrikey string //创建人私钥
}

func (stdtoken *Stdtoken) Address() common.Address {
	address, err := libs.StrKey2Address(stdtoken.CreatorAddressPrikey)
	if err != nil {
		fmt.Printf("StrKey2Address :%v\n", err)
		os.Exit(1)
	}
	return address
}

//获取合约名称合约
func (stdtoken *Stdtoken) Name() (string, error) {

	packed, err := apis.GVA_ABI.Name()
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
	return result, nil
}

func (stdtoken *Stdtoken) Symbol() (string, error) {

	packed, err := apis.GVA_ABI.Symbol()
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

func (stdtoken *Stdtoken) GetDecimals() (*big.Int, error) {

	packed, err := apis.GVA_ABI.GetDecimals()
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

func (stdtoken *Stdtoken) GetTotalSupply() (*big.Int, error) {

	packed, err := apis.GVA_ABI.GetTotalSupply()
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

func (stdtoken *Stdtoken) BalanceOf(account_address string) (*big.Int, error) {

	accoutAddr := common.StrB58ToAddress(account_address)

	cTypeAddr := abi.NewAddress(accoutAddr)
	packed, err := apis.GVA_ABI.BalanceOf(cTypeAddr)
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

func (stdtoken *Stdtoken) Allowance(args reqcontract.StdTokenAllowanceArgs) (*big.Int, error) {

	ownerAddr := common.StrB58ToAddress(args.OwnerAddress)
	cTypeOwnerAddr := abi.NewAddress(ownerAddr)

	addr := common.StrB58ToAddress(args.SpenderAddress)
	cTypeAddr := abi.NewAddress(addr)

	packed, err := apis.GVA_ABI.Allowance(cTypeOwnerAddr, cTypeAddr)
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

func (stdtoken *Stdtoken) Mint(args reqcontract.StdTokenMintArgs) (string, error) {

	address, err := libs.StrKey2Address(args.MintAddress)
	if err != nil {
		return "", fmt.Errorf("invalid MintAddressPriKey to address err:%v", err)
	}
	cTypeAddr := abi.NewAddress(address)
	amount, ok := new(big.Int).SetString(args.Amount, 10)
	if !ok {
		return "", fmt.Errorf("invalid Mint error")
	}
	packed, err := apis.GVA_ABI.Mint(cTypeAddr, abi.NewUint256(amount))
	if err != nil {
		return "", fmt.Errorf("no connection established in service err:%v", err)
	}

	tokenTransfer := new(reqtransfer.StringRawTransaction)
	//初始化GAS和code
	tokenTransfer.To = args.MintAddress
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

func (stdtoken *Stdtoken) Burn(args reqcontract.StdTokenBurnArgs) (string, error) {

	addr := common.StrB58ToAddress(args.BurnAddress)
	cTypeAddr := abi.NewAddress(addr)
	amount, ok := new(big.Int).SetString(args.Amount, 10)
	if !ok {
		return "", fmt.Errorf("invalid Mint error")
	}
	packed, err := apis.GVA_ABI.Burn(cTypeAddr, abi.NewUint256(amount))
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

func (stdtoken *Stdtoken) Approve(args reqcontract.StdTokenApproveArgs) (string, error) {

	addr := common.StrB58ToAddress(args.ApproveSpenderAddress)
	cTypeAddr := abi.NewAddress(addr)
	amount, ok := new(big.Int).SetString(args.Amount, 10)
	if !ok {
		return "", fmt.Errorf("invalid Mint error")
	}
	packed, err := apis.GVA_ABI.Approve(cTypeAddr, abi.NewUint256(amount))
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

func (stdtoken *Stdtoken) TransferFrom(args reqcontract.StdTokenTransferFromArgs) (string, error) {
	address, err := libs.StrKey2Address(args.TransferFromAddress)
	if err != nil {
		return "", fmt.Errorf("invalid TransferFrom FromAddressPreKey err:%v", err)
	}
	cTypeFromAddr := abi.NewAddress(address)

	toAddr := common.StrB58ToAddress(args.TransferToAddress)
	cTypeToAddr := abi.NewAddress(toAddr)
	amount, ok := new(big.Int).SetString(args.TransferFromValue, 10)
	if !ok {
		return "", fmt.Errorf("invalid Mint error")
	}
	packed, err := apis.GVA_ABI.TransferFrom(cTypeFromAddr, cTypeToAddr, abi.NewUint256(amount))
	if err != nil {
		return "", fmt.Errorf("no connection established in service err:%v", err)
	}

	tokenTransfer := new(reqtransfer.StringRawTransaction)
	//初始化GAS和code
	tokenTransfer.To = stdtoken.ContractAddress
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
