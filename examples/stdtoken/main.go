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
	Key              = "0x01010e7d5592d6b4f474ad0b43b2ab77755e9ea6f1cd325e217d78dd57ca9605c636"
	DefaultAddr      = crypto.Prikey2Addr(Key)
	deafaultStdtoken = &stdtoken.StdToken{
		ContractAddress:      "cvsKS3LhJy5sFB62PhBN3HCje8EyzwDpV",
		CreatorAddressPrikey: Key,
	}
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
			Name:     "balanceof",
			Usage:    "<caddress>",
			Category: "arithmetic",
			Action:   Stdtoken_BalanceOf,
		},
		{
			Name:     "caddr",
			Usage:    "<address> <nonce>",
			Category: "arithmetic",
			Action:   Stdtoken_caddr,
		},
		{
			Name:     "approve",
			Usage:    "<spender> <amount> <fromprikey>",
			Category: "arithmetic",
			Action:   Stdtoken_Approve,
		},
		{
			Name:     "transfer",
			Usage:    "<to> <amount>",
			Category: "arithmetic",
			Action:   Stdtoken_Transfer,
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

func Stdtoken_BalanceOf(c *cli.Context) error {
	args := c.Args()
	balance, err := deafaultStdtoken.BalanceOf(args.Get(0))
	if err != nil {
		return err
	}
	fmt.Printf("bal:%v\n", balance.String())
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

func Stdtoken_Transfer(c *cli.Context) error {
	args := c.Args()
	argsTransfer := reqcontract.StdTokenTransferArgs{
		TransferFromAddressPriKey: Key,
		TransferToAddress:         args.Get(0),
		TransferAmount:            args.Get(1),
	}
	txhash, err := deafaultStdtoken.Transfer(argsTransfer)
	if err != nil {
		return err
	}
	fmt.Printf("txhash:%v\n", txhash)
	return nil
}

func Stdtoken_Approve(c *cli.Context) error {
	args := c.Args()

	if c.NArg() < 3 {
		fmt.Println(c.App.Usage)
		return nil
	}
	argsApprove := reqcontract.StdTokenApproveArgs{
		ApproveSpenderAddress:    args.Get(0),
		Amount:                   args.Get(1),
		ApproveFromAddressPriKey: args.Get(2),
	}
	txhash, err := deafaultStdtoken.Approve(argsApprove)
	if err != nil {
		return err
	}
	fmt.Printf("txhash:%v\n", txhash)
	return nil
}
