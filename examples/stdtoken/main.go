package main

// import (
// 	"fmt"
// 	"os"

// 	"github.com/younamebert/xfssdk"
// 	"github.com/younamebert/xfssdk/contract/stdtoken"
// 	reqcontract "github.com/younamebert/xfssdk/servce/contract/request"
// )

// func main() {
// 	key := "0x01010e2f7fdf0c76d4dfaeb76a30aa1a98dcaf5d10dc9596cbafaf5806c14074a813"
// 	handle := xfssdk.Default()
// 	stdtoken := new(stdtoken.StdTokenLocal)
// 	//
// 	argsCreate := reqcontract.TokenArgs{
// 		Name:        "test",
// 		Symbol:      "test",
// 		Decimals:    "10000000",
// 		TotalSupply: "test",
// 	}

// 	code, err := stdtoken.Create(argsCreate)
// 	if err != nil {
// 		fmt.Println(err)
// 		os.Exit(1)
// 		return
// 	}
// 	args := reqcontract.DeployTokenArgs{
// 		Code:       code,
// 		Addresskey: key,
// 	}

// 	stdtoken, txhash, err := stdtoken.DeployToken(args)
// 	if err != nil {
// 		fmt.Println(err)
// 		os.Exit(1)
// 		return
// 	}
// 	fmt.Printf("txhash:%v", txhash)
// 	// stdtoken.Create()
// 	handle.ContractEngine.StdToken = stdtoken
// }
