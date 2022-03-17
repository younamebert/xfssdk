package apistate

import (
	"xfssdk/core/apis"
)

type StateLink interface {
	GetBalance(address string) (*string, error)
	GetBalanceByHash(address string, roothash string) (*string, error)
	GetAccount(address string) (*apis.StateObjResp, error)
	GetAccountByHash(address string, roothash string) (*apis.StateObjResp, error)
}

type ApiState struct {
	// GVA_XFSCLICENT *client.Client
}

// func NewApiState(cli *client.Client) *ApiState {
// 	return &ApiState{
// 		GVA_XFSCLICENT: cli,
// 	}
// }

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
