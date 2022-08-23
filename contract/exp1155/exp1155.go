package exp1155

import (
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

	fmt.Println(address.B58String(), tokenid.String())
	packed, err := apis.GVA_ABI_EXP1155.Mint(abi.NewAddress(address), tokenid)
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
	address := common.StrB58ToAddress(args.Address)

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

func (exp1155 *Exp1155) BalanceOfBatch(accounts, ids []string) (*big.Int, string, error) {
	accoutsCtypes := abi.CTypeString(strings.Join(accounts, ","))
	idsCtypes := abi.CTypeString(strings.Join(ids, ","))

	packed, err := apis.GVA_ABI_EXP1155.BalanceOfBatch(accoutsCtypes, idsCtypes)
	if err != nil {
		return nil, "", fmt.Errorf("no connection established in service err:%v", err)
	}
	req := contract.VMCallData{
		To:   exp1155.ContractAddress,
		Data: packed,
	}
	var result interface{}
	if err := apis.GVA_XFSCLICENT.CallMethod(1, "VM.Call", &req, &result); err != nil {
		return nil, "", err
	}
	if result == nil {
		return nil, "", nil
	}
	tuple, err := abi.DncodeCTypeTuple(result.(string))
	if err != nil {
		return nil, "", err
	}
	bs, _ := json.MarshalIndent(tuple, " ", " ")
	fmt.Println(string(bs))
	// if result == nil {
	// 	return nil, "", nil
	// }
	// tuple, err := abi.DncodeCTypeTuple(result.(string))
	// if err != nil {
	// 	return nil, "", err
	// }
	// newAmount := tuple["CTypeUint256"].(string)
	// amounts, _ := big.NewInt(0).SetString(newAmount, 0)
	// tokenurl := tuple["CTypeString"].(string)
	return nil, "", nil
}

// func (exp1155 *Exp1155) TransferFrom() (string, error) {

// }
