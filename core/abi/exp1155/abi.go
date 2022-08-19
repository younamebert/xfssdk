package exp1155

import (
	"encoding/hex"
	"encoding/json"
	"fmt"

	"github.com/younamebert/xfssdk/common"
)

type Method struct {
	Name       string    `json:"name"`
	Argc       int       `json:"argc"`
	Args       Arguments `json:"args"`
	ReturnType string    `json:"returnType"`
}

type Events struct {
	Name string     `json:"name"`
	Argc int        `json:"argc"`
	Args ArgsEvents `json:"args"`
}
type ArgsEvent struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type ArgsEvents []*ArgsEvent

type ABI struct {
	Events  map[string]*Events
	Methods map[string]*Method
}

func JSON(abi string) (ABI, error) {
	var jsonData ABI
	err := json.Unmarshal([]byte(abi), &jsonData)
	if err != nil {
		return ABI{}, nil
	}
	return jsonData, nil
}

func (abi ABI) PackArgs(name string, args ...interface{}) ([]byte, error) {
	method, exist := abi.Methods[name]
	if !exist {
		return nil, fmt.Errorf("method '%s' not found", name)
	}
	arguments, err := method.Args.Pack(args...)
	if err != nil {
		return nil, err
	}
	return arguments, nil
}

func (abi ABI) Create(params ...interface{}) (string, error) {
	byteArgs, err := abi.PackArgs(CREATE, params...)
	if err != nil {
		return "", err
	}
	data := common.FromHex(EXP1155BIN)
	dataMethod := common.FromHex(CREATE)
	data = append(data, dataMethod...)
	data = append(data, byteArgs...)

	result := "0x" + hex.EncodeToString(data)
	return result, nil
}

func (abi ABI) Mint(params ...interface{}) (string, error) {
	byteArgs, err := abi.PackArgs(MINT, params...)
	if err != nil {
		return "", err
	}
	data := common.FromHex(EXP1155BIN)
	dataMethod := common.FromHex(MINT)
	data = append(data, dataMethod...)
	data = append(data, byteArgs...)

	result := "0x" + hex.EncodeToString(data)
	return result, nil
}

func (abi ABI) MintBatch(params ...interface{}) (string, error) {
	byteArgs, err := abi.PackArgs(MINTBATCH, params...)
	if err != nil {
		return "", err
	}
	data := common.FromHex(EXP1155BIN)
	dataMethod := common.FromHex(MINTBATCH)
	data = append(data, dataMethod...)
	data = append(data, byteArgs...)

	result := "0x" + hex.EncodeToString(data)
	return result, nil
}

func (abi ABI) Burn(params ...interface{}) (string, error) {
	byteArgs, err := abi.PackArgs(BURN, params...)
	if err != nil {
		return "", err
	}
	data := common.FromHex(EXP1155BIN)
	dataMethod := common.FromHex(BURN)
	data = append(data, dataMethod...)
	data = append(data, byteArgs...)

	result := "0x" + hex.EncodeToString(data)
	return result, nil
}

func (abi ABI) IsApprovedForAll(params ...interface{}) (string, error) {
	byteArgs, err := abi.PackArgs(ISAPPROVEDFORALL, params...)
	if err != nil {
		return "", err
	}
	data := common.FromHex(EXP1155BIN)
	dataMethod := common.FromHex(ISAPPROVEDFORALL)
	data = append(data, dataMethod...)
	data = append(data, byteArgs...)

	result := "0x" + hex.EncodeToString(data)
	return result, nil
}

func (abi ABI) SetApprovalForAll(params ...interface{}) (string, error) {
	byteArgs, err := abi.PackArgs(SETAPPROVALFORALL, params...)
	if err != nil {
		return "", err
	}
	data := common.FromHex(EXP1155BIN)
	dataMethod := common.FromHex(SETAPPROVALFORALL)
	data = append(data, dataMethod...)
	data = append(data, byteArgs...)

	result := "0x" + hex.EncodeToString(data)
	return result, nil
}

func (abi ABI) BalanceOfBatch(params ...interface{}) (string, error) {
	byteArgs, err := abi.PackArgs(BALANCEOFBATCH, params...)
	if err != nil {
		return "", err
	}
	data := common.FromHex(EXP1155BIN)
	dataMethod := common.FromHex(BALANCEOFBATCH)
	data = append(data, dataMethod...)
	data = append(data, byteArgs...)

	result := "0x" + hex.EncodeToString(data)
	return result, nil
}

func (abi ABI) BalanceOf(params ...interface{}) (string, error) {
	byteArgs, err := abi.PackArgs(BALANCEOF, params...)
	if err != nil {
		return "", err
	}
	data := common.FromHex(EXP1155BIN)
	dataMethod := common.FromHex(BALANCEOF)
	data = append(data, dataMethod...)
	data = append(data, byteArgs...)

	result := "0x" + hex.EncodeToString(data)
	return result, nil
}
