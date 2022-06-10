package respstate

import "github.com/younamebert/xfssdk/common"

type StateObjResp struct {
	Balance   *string     `json:"balance"`
	Nonce     uint64      `json:"nonce"`
	Extra     *string     `json:"extra"`
	Code      *string     `json:"code"`
	StateRoot common.Hash `json:"state_root"`
}
