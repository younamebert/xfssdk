package stdtoken

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"

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
	fmt.Println(hex.EncodeToString(arguments))
	return arguments, nil
}

func (abi ABI) PackEventsName(events []*Event) ([]*Event, error) {
	var (
		eventsnames        = make(map[string]struct{})
		abieventsnames     = make(map[string]struct{})
		argc           int = len(events)
	)

	for _, obj := range events {
		name := strings.ToLower(obj.Name)
		eventsnames[name] = struct{}{}
	}
	for _, v := range abi.Events {
		if v.Argc == argc {
			for _, obj := range v.Args {
				name := strings.ToLower(obj.Name)
				if _, exist := eventsnames[name]; exist {
					abieventsnames[name] = struct{}{}
				}
			}
		}

		if len(abieventsnames) == len(eventsnames) {
			bs, _ := common.MarshalIndent(v)
			fmt.Println(string(bs))
			return v.Args.Pack(events)
		} else {
			abieventsnames = make(map[string]struct{})
		}
	}
	return nil, fmt.Errorf("not events")
}

func (abi ABI) Create(params ...interface{}) (string, error) {
	args, err := abi.PackArgs(CREATE, params...)
	if err != nil {
		return "", err
	}

	data := common.FromHex(XFSTOKENBin)
	dataMethod := common.FromHex(CREATE)
	data = append(data, dataMethod...)
	data = append(data, args...)

	result := "0x" + hex.EncodeToString(data)
	return result, nil
}

func (abi ABI) Name() (string, error) {
	args, err := abi.PackArgs(NAME)
	if err != nil {
		return "", err
	}

	data := common.FromHex(XFSTOKENBin)
	dataMethod := common.FromHex(NAME)
	data = append(data, dataMethod...)
	data = append(data, args...)

	result := "0x" + hex.EncodeToString(data)
	return result, nil
}

func (abi ABI) Symbol() (string, error) {
	args, err := abi.PackArgs(SYMBOL)
	if err != nil {
		return "", err
	}

	data := common.FromHex(XFSTOKENBin)
	dataMethod := common.FromHex(SYMBOL)
	data = append(data, dataMethod...)
	data = append(data, args...)

	result := "0x" + hex.EncodeToString(data)
	return result, nil
}

func (abi ABI) GetDecimals() (string, error) {
	args, err := abi.PackArgs(GETDECIMALS)
	if err != nil {
		return "", err
	}

	data := common.FromHex(XFSTOKENBin)
	dataMethod := common.FromHex(GETDECIMALS)
	data = append(data, dataMethod...)
	data = append(data, args...)

	result := "0x" + hex.EncodeToString(data)
	return result, nil
}

func (abi ABI) GetTotalSupply() (string, error) {
	args, err := abi.PackArgs(GETTOTALSUPPLY)
	if err != nil {
		return "", err
	}

	data := common.FromHex(XFSTOKENBin)
	dataMethod := common.FromHex(GETTOTALSUPPLY)
	data = append(data, dataMethod...)
	data = append(data, args...)

	result := "0x" + hex.EncodeToString(data)
	return result, nil
}

func (abi ABI) BalanceOf(params ...interface{}) (string, error) {
	byteArgs, err := abi.PackArgs(BALANCEOF, params...)
	if err != nil {
		return "", err
	}

	data := common.FromHex(XFSTOKENBin)
	dataMethod := common.FromHex(BALANCEOF)
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

	data := common.FromHex(XFSTOKENBin)
	dataMethod := common.FromHex(MINT)
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

	data := common.FromHex(XFSTOKENBin)
	dataMethod := common.FromHex(BURN)
	data = append(data, dataMethod...)
	data = append(data, byteArgs...)

	result := "0x" + hex.EncodeToString(data)
	return result, nil
}

func (abi ABI) Approve(params ...interface{}) (string, error) {
	byteArgs, err := abi.PackArgs(APPROVE, params...)
	if err != nil {
		return "", err
	}

	data := common.FromHex(XFSTOKENBin)
	dataMethod := common.FromHex(APPROVE)
	data = append(data, dataMethod...)
	data = append(data, byteArgs...)

	result := "0x" + hex.EncodeToString(data)
	return result, nil
}

func (abi ABI) Allowance(params ...interface{}) (string, error) {
	byteArgs, err := abi.PackArgs(ALLOWANCE, params...)
	if err != nil {
		return "", err
	}

	data := common.FromHex(XFSTOKENBin)
	dataMethod := common.FromHex(ALLOWANCE)
	data = append(data, dataMethod...)
	data = append(data, byteArgs...)

	result := "0x" + hex.EncodeToString(data)
	return result, nil
}

func (abi ABI) TransferFrom(params ...interface{}) (string, error) {
	byteArgs, err := abi.PackArgs(TRANSFERFROM, params...)
	if err != nil {
		return "", err
	}

	data := common.FromHex(XFSTOKENBin)
	dataMethod := common.FromHex(TRANSFERFROM)
	data = append(data, dataMethod...)
	data = append(data, byteArgs...)

	result := "0x" + hex.EncodeToString(data)
	return result, nil
}
