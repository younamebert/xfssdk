package apinet

import "github.com/younamebert/xfssdk/core/apis"

type NetLink interface {
	GetNodeId() (*string, error)
}

type ApiNet struct {
	// GVA_XFSCLICENT *client.Client
}

// func (net *ApiNet) GetPeers() ([]string, error) {
// 	result := make([]string, 0)
// 	if err := apis.GVA_XFSCLICENT.CallMethod(1, "Net.GetPeers", nil, &result); err != nil {
// 		return nil, err
// 	}
// 	return result, nil
// }

// func NewApiNet(cli *client.Client) *ApiNet {
// 	return &ApiNet{
// 		GVA_XFSCLICENT: cli,
// 	}
// }

func (net *ApiNet) GetNodeId() (*string, error) {
	var result *string
	if err := apis.GVA_XFSCLICENT.CallMethod(1, "Net.GetPeers", nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}
