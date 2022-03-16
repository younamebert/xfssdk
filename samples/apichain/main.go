package main

import (
	"fmt"
	"os"
	"xfssdk/api"
	"xfssdk/common"
	"xfssdk/core/apis"
	"xfssdk/libs/client"
)

func main() {
	cli := client.NewClient("https://api.scan.xfs.tech/jsonrpc/v2/", "5s")
	apis.SetXFSClient(cli)

	apimethod := api.NewApiMethod()
	latestBlockHeader, err := apimethod.Chain.GetHead()
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
