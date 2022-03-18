package main

import (
	"fmt"
	"os"

	"github.com/younamebert/xfssdk/api"
	"github.com/younamebert/xfssdk/common"
	"github.com/younamebert/xfssdk/core/apis"
	"github.com/younamebert/xfssdk/libs/client"
)

func main() {
	cli := client.NewClient("https://api.scan.xfs.tech/jsonrpc/v2/", "5s")
	apis.SetXFSClient(cli)

	apimethod := api.NewApiMethod()
	latestBlockHeader, err := apimethod.Chain.GetBlockByNumber("11")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	bs, err := common.MarshalIndent(latestBlockHeader)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(string(bs))
}
