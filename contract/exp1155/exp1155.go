package exp1155

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"strings"

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

type Exp1155Local struct{}

func (exp1155local *Exp1155Local) Create(args *reqcontract.Exp1155CreateArgs) (string, error) {
	code, err := apis.GVA_ABI_EXP1155.Create()
	if err != nil {
		return "", fmt.Errorf("an exception occurred of contract argument err:%v", err)
	}
	return code, err
}

func (exp1155local *Exp1155Local) Deploy(args reqcontract.Exp1155DeployArgs) (*Exp1155, string, error) {
	//创建合约交易
	tokenTransfer := new(reqtransfer.StringRawTransaction)
	tokenTransfer.Data = args.Code

	address, err := libs.StrKey2Address(args.Privkey)
	if err != nil {
		return nil, "", err
	}
	//签名交易
	stdtokentransfer, err := transfer.EnCodeRawTransaction(args.Privkey, tokenTransfer)
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
	result := &Exp1155{
		ContractAddress:      caddr.B58String(), //合约地址
		CreatorAddressPrikey: args.Privkey,      //创建人私钥
	}
	return result, txhash, nil
}

type Exp1155 struct {
	ContractAddress      string //合约地址
	CreatorAddressPrikey string //创建人私钥
}

func (exp1155 *Exp1155) Mint(args *reqcontract.Exp1155MintArgs) (string, error) {

	address := common.StrB58ToAddress(args.Recipient)
	tokenid := abi.CTypeString(args.Tokenurl)
	amount := abi.NewUint256(args.Amount)

	packed, err := apis.GVA_ABI_EXP1155.Mint(abi.NewAddress(address), tokenid, amount)
	if err != nil {
		return "", fmt.Errorf("no connection established in service err:%v", err)
	}

	from, err := libs.StrKey2Address(exp1155.CreatorAddressPrikey)
	if err != nil {
		return "", err
	}
	req := contract.VMCallData{
		From: from.B58String(),
		To:   exp1155.ContractAddress,
		Data: packed,
	}
	var result interface{}
	if err := apis.GVA_XFSCLICENT.CallMethod(1, "VM.Call", &req, &result); err != nil {
		return "", err
	}
	tokenTransfer := new(reqtransfer.StringRawTransaction)
	//初始化GAS和code
	tokenTransfer.To = exp1155.ContractAddress
	tokenTransfer.Data = packed
	// fmt.Printf()
	stdtokentransfer, err := transfer.EnCodeRawTransaction(exp1155.CreatorAddressPrikey, tokenTransfer)
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

func (exp1155 *Exp1155) MintBatch(args *reqcontract.Exp1155MintBatchArgs) (string, error) {
	address := common.StrB58ToAddress(args.Recipient)

	if len(args.Amounts) != len(args.TokenUrls) {
		return "", errors.New("Amounts.Len != tokenurls.Len")
	}
	var amounts, tokenUrls abi.CTypeString
	amounts = abi.CTypeString(strings.Join(args.Amounts, ","))
	tokenUrls = abi.CTypeString(strings.Join(args.TokenUrls, ","))
	packed, err := apis.GVA_ABI_EXP1155.MintBatch(abi.NewAddress(address), tokenUrls, amounts)
	if err != nil {
		return "", fmt.Errorf("no connection established in service err:%v", err)
	}
	from, err := libs.StrKey2Address(exp1155.CreatorAddressPrikey)
	if err != nil {
		return "", err
	}
	req := contract.VMCallData{
		From: from.B58String(),
		To:   exp1155.ContractAddress,
		Data: packed,
	}
	var result []byte
	if err := apis.GVA_XFSCLICENT.CallMethod(1, "VM.Call", &req, &result); err != nil {
		return "", err
	}
	tokenTransfer := new(reqtransfer.StringRawTransaction)
	//初始化GAS和code
	tokenTransfer.To = exp1155.ContractAddress
	tokenTransfer.Data = packed
	// fmt.Printf()
	stdtokentransfer, err := transfer.EnCodeRawTransaction(exp1155.CreatorAddressPrikey, tokenTransfer)
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

func (exp1155 *Exp1155) BalanceOf(account_address string, tokenid *big.Int) (string, error) {
	accoutAddr := common.StrB58ToAddress(account_address)
	cTypeAddr := abi.NewAddress(accoutAddr)
	id := abi.NewUint256(tokenid)
	packed, err := apis.GVA_ABI_EXP1155.BalanceOf(cTypeAddr, id)
	if err != nil {
		return "", fmt.Errorf("no connection established in service err:%v", err)
	}
	req := contract.VMCallData{
		To:   exp1155.ContractAddress,
		Data: packed,
	}
	var result interface{}
	if err := apis.GVA_XFSCLICENT.CallMethod(1, "VM.Call", &req, &result); err != nil {
		return "", err
	}
	if result == nil {
		return "", nil
	}
	tuple, err := abi.DncodeCTypeTuple(result.(string))
	if err != nil {
		return "", err
	}
	bs, err := json.Marshal(tuple)
	if err != nil {
		return "", err
	}
	return string(bs), nil
}

func (exp1155 *Exp1155) BalanceOfBatch(accounts, ids []string) (string, error) {
	accoutsCtypes := abi.CTypeString(strings.Join(accounts, ","))
	idsCtypes := abi.CTypeString(strings.Join(ids, ","))

	packed, err := apis.GVA_ABI_EXP1155.BalanceOfBatch(accoutsCtypes, idsCtypes)
	if err != nil {
		return "", fmt.Errorf("no connection established in service err:%v", err)
	}
	req := contract.VMCallData{
		To:   exp1155.ContractAddress,
		Data: packed,
	}
	var result interface{}
	if err := apis.GVA_XFSCLICENT.CallMethod(1, "VM.Call", &req, &result); err != nil {
		return "", err
	}
	if result == nil {
		return "", nil
	}
	tuple, err := abi.DncodeCTypeTuple(result.(string))
	if err != nil {
		return "", err
	}
	bs, err := json.Marshal(tuple)
	if err != nil {
		return "", err
	}
	return string(bs), nil
}

func (exp1155 *Exp1155) SetApprovalForAll(args *reqcontract.Exp1155SetApprovalForAllArgs) (string, error) {
	operator := common.StrB58ToAddress(args.Operator)

	var (
		approvedCType abi.CTypeBool
		operatorCType abi.CTypeAddress
	)
	if args.Approved {
		approvedCType = abi.CTypeBool{1}
	} else {
		approvedCType = abi.CTypeBool{0}
	}
	operatorCType = abi.NewAddress(operator)

	packed, err := apis.GVA_ABI_EXP1155.SetApprovalForAll(operatorCType, approvedCType)
	if err != nil {
		return "", fmt.Errorf("no connection established in service err:%v", err)
	}
	from, err := libs.StrKey2Address(args.PriKey)
	if err != nil {
		return "", err
	}
	req := contract.VMCallData{
		From: from.B58String(),
		To:   exp1155.ContractAddress,
		Data: packed,
	}
	var result interface{}
	if err := apis.GVA_XFSCLICENT.CallMethod(1, "VM.Call", &req, &result); err != nil {
		return "", err
	}
	fmt.Println(result)
	tokenTransfer := new(reqtransfer.StringRawTransaction)
	//初始化GAS和code
	tokenTransfer.To = exp1155.ContractAddress
	tokenTransfer.Data = packed
	// fmt.Printf()
	stdtokentransfer, err := transfer.EnCodeRawTransaction(args.PriKey, tokenTransfer)
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

func (exp1155 *Exp1155) IsApprovalForAll(args *reqcontract.Exp1155IsApprovedForAllArgs) (bool, error) {
	operator, account := common.StrB58ToAddress(args.Operator), common.StrB58ToAddress(args.Account)

	operatorCType, accountCType := abi.NewAddress(operator), abi.NewAddress(account)

	packed, err := apis.GVA_ABI_EXP1155.IsApprovedForAll(accountCType, operatorCType)
	if err != nil {
		return false, fmt.Errorf("no connection established in service err:%v", err)
	}
	req := contract.VMCallData{
		To:   exp1155.ContractAddress,
		Data: packed,
	}

	var result string
	if err := apis.GVA_XFSCLICENT.CallMethod(1, "VM.Call", &req, &result); err != nil {
		return false, err
	}
	result = strings.TrimPrefix(result, "0x")
	bytesBool, err := hex.DecodeString(result)
	if err != nil {
		return false, err
	}
	CTypeBool := abi.CTypeBool{bytesBool[0]}
	return CTypeBool.Bool(), nil
}

func (exp1155 *Exp1155) TransferFrom(args *reqcontract.Exp1155SafeTransferFromArgs) (string, error) {
	from, to := common.StrB58ToAddress(args.From), common.StrB58ToAddress(args.To)
	amountCtype, tokenidCtype := abi.NewUint256(args.Amount), abi.NewUint256(args.Id)
	fromCtype, toCtype := abi.NewAddress(from), abi.NewAddress(to)

	packed, err := apis.GVA_ABI_EXP1155.TransferFrom(fromCtype, toCtype, tokenidCtype, amountCtype)
	if err != nil {
		return "", fmt.Errorf("no connection established in service err:%v", err)
	}

	caller, err := libs.StrKey2Address(args.Prikey)
	if err != nil {
		return "", err
	}
	req := contract.VMCallData{
		From: caller.B58String(),
		To:   exp1155.ContractAddress,
		Data: packed,
	}
	var result string
	if err := apis.GVA_XFSCLICENT.CallMethod(1, "VM.Call", &req, &result); err != nil {
		return "", err
	}
	tokenTransfer := new(reqtransfer.StringRawTransaction)
	//初始化GAS和code
	tokenTransfer.To = exp1155.ContractAddress
	tokenTransfer.Data = packed

	stdtokentransfer, err := transfer.EnCodeRawTransaction(args.Prikey, tokenTransfer)
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

func (exp1155 *Exp1155) TransferFromBatch(args *reqcontract.Exp1155SafeBatchTransferFromArgs) (string, error) {
	from, to := common.StrB58ToAddress(args.From), common.StrB58ToAddress(args.To)

	fromCtype, toCtype := abi.NewAddress(from), abi.NewAddress(to)

	amountsCtype, tokenidsCtype := abi.CTypeString(args.Amounts), abi.CTypeString(args.Ids)
	packed, err := apis.GVA_ABI_EXP1155.TransferFromBatch(fromCtype, toCtype, tokenidsCtype, amountsCtype)
	if err != nil {
		return "", fmt.Errorf("no connection established in service err:%v", err)
	}
	caller, err := libs.StrKey2Address(args.Prikey)
	if err != nil {
		return "", err
	}
	req := contract.VMCallData{
		From: caller.B58String(),
		To:   exp1155.ContractAddress,
		Data: packed,
	}
	var result string
	if err := apis.GVA_XFSCLICENT.CallMethod(1, "VM.Call", &req, &result); err != nil {
		return "", err
	}
	tokenTransfer := new(reqtransfer.StringRawTransaction)
	//初始化GAS和code
	tokenTransfer.To = exp1155.ContractAddress
	tokenTransfer.Data = packed

	stdtokentransfer, err := transfer.EnCodeRawTransaction(args.Prikey, tokenTransfer)
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

func (exp1155 *Exp1155) Burn(args *reqcontract.Exp1155BurnArgs) (string, error) {
	address := common.StrB58ToAddress(args.From)
	fromCtype := abi.NewAddress(address)
	tokenidCtype := abi.NewUint256(args.ID)
	amountCtype := abi.NewUint256(args.Amount)

	packed, err := apis.GVA_ABI_EXP1155.Burn(fromCtype, tokenidCtype, amountCtype)
	if err != nil {
		return "", fmt.Errorf("no connection established in service err:%v", err)
	}
	caller, err := libs.StrKey2Address(args.Prikey)
	if err != nil {
		return "", err
	}
	req := contract.VMCallData{
		From: caller.B58String(),
		To:   exp1155.ContractAddress,
		Data: packed,
	}
	var result string
	if err := apis.GVA_XFSCLICENT.CallMethod(1, "VM.Call", &req, &result); err != nil {
		return "", err
	}
	tokenTransfer := new(reqtransfer.StringRawTransaction)
	//初始化GAS和code
	tokenTransfer.To = exp1155.ContractAddress
	tokenTransfer.Data = packed

	stdtokentransfer, err := transfer.EnCodeRawTransaction(args.Prikey, tokenTransfer)
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

func (exp1155 *Exp1155) BurnBatch(args *reqcontract.Exp1155BurnBatchArgs) (string, error) {
	address := common.StrB58ToAddress(args.From)
	fromCtype := abi.NewAddress(address)
	tokenidCtype, amountsCtype := abi.CTypeString(args.IDs), abi.CTypeString(args.Amounts)

	packed, err := apis.GVA_ABI_EXP1155.BurnBatch(fromCtype, tokenidCtype, amountsCtype)
	if err != nil {
		return "", fmt.Errorf("no connection established in service err:%v", err)
	}
	caller, err := libs.StrKey2Address(args.Prikey)
	if err != nil {
		return "", err
	}
	req := contract.VMCallData{
		From: caller.B58String(),
		To:   exp1155.ContractAddress,
		Data: packed,
	}
	var result string
	if err := apis.GVA_XFSCLICENT.CallMethod(1, "VM.Call", &req, &result); err != nil {
		return "", err
	}
	tokenTransfer := new(reqtransfer.StringRawTransaction)
	//初始化GAS和code
	tokenTransfer.To = exp1155.ContractAddress
	tokenTransfer.Data = packed

	stdtokentransfer, err := transfer.EnCodeRawTransaction(args.Prikey, tokenTransfer)
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
