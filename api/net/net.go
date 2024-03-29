package apinet

import "github.com/younamebert/xfssdk/core/apis"

type NetLink interface {
	GetNodeId() (*string, error)
}

type ApiNet struct{}

// GetNodeId get nodeid of current node chain
func (net *ApiNet) GetNodeId() (*string, error) {
	var result *string
	if err := apis.GVA_XFSCLICENT.CallMethod(1, "Net.GetPeers", nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}
