package apistate

import (
	"github.com/younamebert/xfssdk/core/apis"
)

type StateLink interface {
	GetBalance(address string) (*string, error)
	GetBalanceByHash(address string, roothash string) (*string, error)
	GetAccount(address string) (*apis.StateObjResp, error)
	GetAccountByHash(address string, roothash string) (*apis.StateObjResp, error)
}

// ApiState 账户状态结构体
type ApiState struct{}

// GetBalance 指定账户地址查看账户余额
func (state *ApiState) GetBalance(address string) (*string, error) {
	req := &GetBalanceArgs{
		Address: address,
	}
	var result *string
	if err := apis.GVA_XFSCLICENT.CallMethod(1, "State.GetBalance", &req, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetBalanceByHash 指定账户地址和区块hash查看账户余额
func (state *ApiState) GetBalanceByHash(address string, roothash string) (*string, error) {
	req := &GetBalanceArgs{
		Address:  address,
		RootHash: roothash,
	}
	var result *string
	if err := apis.GVA_XFSCLICENT.CallMethod(1, "State.GetBalance", &req, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetAccount 指定账户地址获取账户信息
func (state *ApiState) GetAccount(address string) (*apis.StateObjResp, error) {
	req := &GetAccountArgs{
		Address: address,
	}
	result := new(apis.StateObjResp)
	if err := apis.GVA_XFSCLICENT.CallMethod(1, "State.GetAccount", &req, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetAccountByHash 指定账户地址和区块hash查看账户信息
func (state *ApiState) GetAccountByHash(address string, roothash string) (*apis.StateObjResp, error) {
	req := &GetAccountArgs{
		Address:  address,
		RootHash: roothash,
	}
	result := new(apis.StateObjResp)
	if err := apis.GVA_XFSCLICENT.CallMethod(1, "State.GetAccount", &req, &result); err != nil {
		return nil, err
	}
	return result, nil
}
