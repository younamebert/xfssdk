package main

import (
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"
	"strings"

	"github.com/younamebert/xfssdk"
	"github.com/younamebert/xfssdk/common"
	"github.com/younamebert/xfssdk/contract/exp1155"
	"github.com/younamebert/xfssdk/crypto"
	reqcontract "github.com/younamebert/xfssdk/servce/contract/request"
	"gopkg.in/urfave/cli.v1"
)

var (
	Key              = "0x0101da503ac2fe8afa56ab4f6ac3443c1c8051e02d67bd7670c8d86a5e9f42c8c58d"
	DefaultAddr      = crypto.Prikey2Addr(Key)
	deafaultStdtoken = &exp1155.Exp1155{
		ContractAddress:      "bZBJwCVZNvtnZ6MQRhoDy54EJPzTo8czc",
		CreatorAddressPrikey: Key,
	}
	app          *cli.App
	handle       = xfssdk.Default()
	exp1155Local = new(exp1155.Exp1155Local)
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
			Usage:    "create",
			Category: "arithmetic",
			Action:   exp1155Token_Create,
		},
		{
			Name:     "deploy",
			Usage:    "deploy <code> <addrprikey>",
			Category: "arithmetic",
			Action:   exp1155Token_Deploy,
		},
		{
			Name:     "mint",
			Usage:    "<address> <tokenurl>",
			Category: "arithmetic",
			Action:   exp1155Token_Mint,
		},
		{
			Name:     "mintBatch",
			Usage:    "<address> <amount> <tokenUri>",
			Category: "arithmetic",
			Action:   exp1155Token_MintBatch,
		},
		{
			Name:     "balanceof",
			Usage:    "<address> <tokenid>",
			Category: "arithmetic",
			Action:   exp1155Token_BalanceOf,
		},
		{
			Name:     "caddr",
			Usage:    "<address> <nonce>",
			Category: "arithmetic",
			Action:   Stdtoken_caddr,
		},
		// {
		// 	Name:     "approve",
		// 	Usage:    "<spender> <amount> <fromprikey>",
		// 	Category: "arithmetic",
		// 	Action:   Stdtoken_Approve,
		// },
		// {
		// 	Name:     "transfer",
		// 	Usage:    "<to> <amount>",
		// 	Category: "arithmetic",
		// 	Action:   Stdtoken_Transfer,
		// },
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func exp1155Token_Create(c *cli.Context) error {
	argsCreate := new(reqcontract.Exp1155CreateArgs)
	code, err := exp1155Local.Create(argsCreate)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return err
	}
	fmt.Println(code)
	return nil
}

func exp1155Token_Deploy(c *cli.Context) error {
	args := c.Args()

	if c.NArg() < 1 {
		fmt.Println(c.App.Usage)
		return nil
	}

	argsDeploy := reqcontract.Exp1155DeployArgs{
		Code:    args.Get(0),
		Privkey: Key,
	}
	_, txhash, err := exp1155Local.Deploy(argsDeploy)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return err
	}
	fmt.Println(txhash)
	return nil
}

func exp1155Token_Mint(c *cli.Context) error {
	args := c.Args()

	if c.NArg() < 2 {
		fmt.Println(c.App.Usage)
		return nil
	}
	Exp1155Class := &reqcontract.Exp1155MintArgs{
		Recipient: args.Get(0),
		Tokenurl:  args.Get(1),
	}

	txhash, err := deafaultStdtoken.Mint(Exp1155Class)
	if err != nil {
		return err
	}
	fmt.Println(txhash)
	return nil
}

func exp1155Token_MintBatch(c *cli.Context) error {
	args := c.Args()
	// amount, err := strconv.Atoi(args.Get(1))
	// if err != nil {
	// 	return err
	// }
	tokenurls := strings.Split(args.Get(1), ",")
	Exp1155Class := &reqcontract.NFTMarketMintBatchArgs{
		Address:   args.Get(0),
		Amount:    big.NewInt(int64(amount)),
		TokenUrls: tokenurls,
	}

	txhash, err := deafaultStdtoken.Mint(Exp1155Class)
	if err != nil {
		return err
	}
	fmt.Println(txhash)
	return nil
}
func exp1155Token_BalanceOf(c *cli.Context) error {
	args := c.Args()

	id, err := strconv.Atoi(args.Get(1))
	if err != nil {
		return err
	}
	tokenid := big.NewInt(int64(id))
	txhash, err := deafaultStdtoken.BalanceOf(args.Get(0), tokenid)
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
