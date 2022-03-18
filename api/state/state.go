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

type ApiState struct{}

// GetBalance specify an account address to view the account balance
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

// GetBalanceByHash specify the account address and block hash to view the account balance
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

// GetAccount specify an account address to obtain account information
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

// GetAccountByHash specify account address and block hash to view account information
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
