package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/younamebert/xfssdk"
	"github.com/younamebert/xfssdk/common"
	"github.com/younamebert/xfssdk/common/ahash"
	"github.com/younamebert/xfssdk/contract/stdtoken"
	"github.com/younamebert/xfssdk/crypto"
	"github.com/younamebert/xfssdk/libs"
	reqcontract "github.com/younamebert/xfssdk/servce/contract/request"
	"gopkg.in/urfave/cli.v1"
)

var (
	Key           = "0x01012ad0c20731aa20999e5424e09329eaf421aca466e10aaa5cfc3ccc268aefc2aa"
	app           *cli.App
	handle        = xfssdk.Default()
	stdtokenLocal = new(stdtoken.StdTokenLocal)
)

// var app = cli.NewApp()

func init() {
	app = cli.NewApp()
	app.Name = "xfssdk"
	app.Usage = "xfssdk stdtoken"
	app.Version = "1.0.0"
	// handle.ContractEngine.StdToken = new(stdtoken.StdTokenLocal)
}

func main() {
	app.Commands = []cli.Command{
		{
			Name: "create",
			// Aliases:  []string{"create"},
			Usage:    "create <name> <symbol> <decimals> <totalSupply>",
			Category: "arithmetic",
			Action:   Stdtoken_Create,
		},
		{
			Name:     "deploy",
			Usage:    "deploy <code> <addrprikey>",
			Category: "arithmetic",
			Action:   Stdtoken_Deploy,
		},
		{
			Name:     "mint",
			Usage:    "<addrprikey> <nonce> <amount>",
			Category: "arithmetic",
			Action:   Stdtoken_Mint,
		},
		{
			Name:     "caddr",
			Usage:    "<address> <nonce>",
			Category: "arithmetic",
			Action:   Stdtoken_caddr,
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func Stdtoken_Create(c *cli.Context) error {
	args := c.Args()

	if c.NArg() < 4 {
		fmt.Println(c.App.Usage)
		return nil
	}
	argsCreate := reqcontract.TokenArgs{
		Name:        args.Get(0),
		Symbol:      args.Get(1),
		Decimals:    args.Get(2),
		TotalSupply: args.Get(3),
	}

	code, err := stdtokenLocal.Create(argsCreate)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return err
	}
	fmt.Println(code)
	return nil
}

func Stdtoken_Deploy(c *cli.Context) error {
	args := c.Args()

	if c.NArg() < 1 {
		fmt.Println(c.App.Usage)
		return nil
	}

	argsDeploy := reqcontract.DeployTokenArgs{
		Code:       args.Get(0),
		Addresskey: Key,
	}
	_, txhash, err := stdtokenLocal.DeployToken(argsDeploy)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return err
	}
	fmt.Println(txhash)
	return nil
}

func Stdtoken_Mint(c *cli.Context) error {
	args := c.Args()

	addr, err := libs.StrKey2Address(args.Get(0))
	if err != nil {
		return err
	}
	nonceInt, err := strconv.Atoi(args.Get(1))
	if err != nil {
		return err
	}
	fromAddressHashBytes := ahash.SHA256(addr[:])
	fromAddressHash := common.Bytes2Hash(fromAddressHashBytes)
	caddr := crypto.CreateAddress(fromAddressHash, uint64(nonceInt))
	stdtokenClass := &stdtoken.StdToken{
		ContractAddress:      caddr.B58String(),
		CreatorAddressPrikey: args.Get(0),
	}

	argsMint := reqcontract.StdTokenMintArgs{
		Amount: args.Get(2),
	}
	txhash, err := stdtokenClass.Mint(argsMint)
	if err != nil {
		return err
	}
	fmt.Println(txhash)
	return nil
}

func Stdtoken_caddr(c *cli.Context) error {
	args := c.Args()
	address := common.StrB58ToAddress(args.Get(0))
	nonceInt, err := strconv.Atoi(args.Get(1))
	if err != nil {
		return err
	}
	caddr := crypto.GetCAddr(address, uint64(nonceInt))
	fmt.Println(caddr.B58String())
	return nil
}
