package client

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
)

type Client struct {
	hostUrl string
	timeOut string
}

type jsonRPCReq struct {
	JsonRPC string      `json:"jsonrpc"`
	ID      int         `json:"id"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
}

type jsonRPCResp struct {
	JSONRPC string      `json:"jsonrpc"`
	Result  interface{} `json:"result"`
	Error   rpcError    `json:"error"`
	ID      int         `json:"id"`
}

func NewClient(url, timeOut string) *Client {
	return &Client{
		hostUrl: url,
		timeOut: timeOut,
	}
}

// CallMethod executes a JSON-RPC call with the given psrameters,which is important to the rpc server.
func (cli *Client) CallMethod(id int, methodname string, params interface{}, out interface{}) error {
	client := resty.New()

	timeDur, err := time.ParseDuration(cli.timeOut)
	if err != nil {
		return err
	}
	client = client.SetTimeout(timeDur)
	req := &jsonRPCReq{
		JsonRPC: "2.0",
		ID:      id,
		Method:  methodname,
		Params:  params,
	}
	// The result must be a pointer so that response json can unmarshal into it.
	var resp *jsonRPCResp = nil
	r, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(req).
		SetResult(&resp). // or SetResult(AuthSuccess{}).
		Post(cli.hostUrl)
	if err != nil {
		return err
	}
	if resp == nil {
		return nil
	}
	e := resp.Error.Message
	if e != "" {
		return fmt.Errorf(e)
	}

	js, err := json.Marshal(resp.Result)
	if err != nil {
		return err
	}
	err = json.Unmarshal(js, out)
	if err != nil {
		return err
	}
	_ = r
	return nil
}
