package apinet

import "xfssdk/core/apis"

type NetLink interface {
	GetNodeId() (*string, error)
}

type ApiNet struct {
	// XFSCLICENT *client.Client
}

// func (net *ApiNet) GetPeers() ([]string, error) {
// 	result := make([]string, 0)
// 	if err := apis.XFSCLICENT.CallMethod(1, "Net.GetPeers", nil, &result); err != nil {
// 		return nil, err
// 	}
// 	return result, nil
// }

// func NewApiNet(cli *client.Client) *ApiNet {
// 	return &ApiNet{
// 		XFSCLICENT: cli,
// 	}
// }

func (net *ApiNet) GetNodeId() (*string, error) {
	var result *string
	if err := apis.XFSCLICENT.CallMethod(1, "Net.GetPeers", nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}
