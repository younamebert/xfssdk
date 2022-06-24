package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/younamebert/xfssdk"
	"github.com/younamebert/xfssdk/common"
	"github.com/younamebert/xfssdk/common/ahash"
	"github.com/younamebert/xfssdk/contract/nfttoken"
	"github.com/younamebert/xfssdk/crypto"
	"github.com/younamebert/xfssdk/libs"
	reqcontract "github.com/younamebert/xfssdk/servce/contract/request"
	"gopkg.in/urfave/cli.v1"
)

var (
	Key          = "0x010140e13c81c518c4452048eeeef601dafd66d1621be9ee696c71c1916440be9aa3"
	app          *cli.App
	handle       = xfssdk.Default()
	nftokenLocal = new(nfttoken.NFTTokenLocal)
)

// var app = cli.NewApp()

func init() {
	app = cli.NewApp()
	app.Name = "xfssdk"
	app.Usage = "xfssdk nftoken"
	app.Version = "1.0.0"
	// handle.ContractEngine.StdToken = new(stdtoken.StdTokenLocal)
}

func main() {
	app.Commands = []cli.Command{
		{
			Name:     "create",
			Usage:    "create <name> <symbol>",
			Category: "arithmetic",
			Action:   NFToken_Create,
		},
		{
			Name:     "deploy",
			Usage:    "deploy <code> <addrprikey>",
			Category: "arithmetic",
			Action:   NFToken_Deploy,
		},
		{
			Name:     "mint",
			Usage:    "<addrprikey> <nonce> <tokenUri>",
			Category: "arithmetic",
			Action:   NFToken_Mint,
		},
		{
			Name:     "caddr",
			Usage:    "<address> <nonce>",
			Category: "arithmetic",
			Action:   NFToken_caddr,
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func NFToken_Create(c *cli.Context) error {
	args := c.Args()

	if c.NArg() < 2 {
		fmt.Println(c.App.Usage)
		return nil
	}
	argsCreate := reqcontract.NFTTokenCreateArgs{
		Name:   args.Get(0),
		Symbol: args.Get(1),
	}

	code, err := nftokenLocal.NFTCreate(argsCreate)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return err
	}
	fmt.Println(code)
	return nil
}

func NFToken_Deploy(c *cli.Context) error {
	args := c.Args()

	if c.NArg() < 1 {
		fmt.Println(c.App.Usage)
		return nil
	}

	argsDeploy := reqcontract.DeployNFTokenArgs{
		Code:       args.Get(0),
		Addresskey: Key,
	}
	_, txhash, err := nftokenLocal.NFTDeployToken(argsDeploy)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return err
	}
	fmt.Println(txhash)
	return nil
}

func NFToken_Mint(c *cli.Context) error {
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
	nftokenClass := &nfttoken.NFToken{
		ContractAddress:      caddr.B58String(),
		CreatorAddressPrikey: args.Get(0),
	}

	argsMint := reqcontract.NFTokenMintArgs{
		TokenId: args.Get(2),
	}
	txhash, err := nftokenClass.Mint(argsMint)
	if err != nil {
		return err
	}
	fmt.Println(txhash)
	return nil
}

func NFToken_caddr(c *cli.Context) error {
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
