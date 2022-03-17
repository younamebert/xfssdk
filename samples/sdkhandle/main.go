package main

import (
	"fmt"
	"os"

	"github.com/younamebert/xfssdk"
	"github.com/younamebert/xfssdk/common"
)

func main() {
	handle := xfssdk.Default()
	latestBlockHeader, err := handle.ApiMethod.Chain.GetHead()
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
