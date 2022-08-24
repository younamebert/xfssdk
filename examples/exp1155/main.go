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
	Key                 = "0x0101da503ac2fe8afa56ab4f6ac3443c1c8051e02d67bd7670c8d86a5e9f42c8c58d"
	DefaultAddr         = crypto.Prikey2Addr(Key)
	defaultEXP1155Token = &exp1155.Exp1155{
		ContractAddress:      "ddjWAUubn6JA4tu51Y1vyjxSNyGVpr3jD",
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
			Usage:    "<address> <amounts> <tokenUrls>",
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
			Name:     "balanceofBatch",
			Usage:    "<addrs> <tokenids>",
			Category: "arithmetic",
			Action:   exp1155Token_BalanceOfBatch,
		},
		{
			Name:     "approvalforall",
			Usage:    "<operator> <approved> <prikey>",
			Category: "arithmetic",
			Action:   exp1155Token_SetApprovalForAll,
		},
		{
			Name:     "isApprovalforall",
			Usage:    "<account> <operator>",
			Category: "arithmetic",
			Action:   exp1155Token_IsApprovalForAll,
		},
		{
			Name:     "transferFrom",
			Usage:    "<from> <to> <id> <amount> <prikey>",
			Category: "arithmetic",
			Action:   exp1155Token_SafeTransferFrom,
		},
		{
			Name:     "transferFromBatch",
			Usage:    "<from> <to> <ids> <amounts>",
			Category: "arithmetic",
			Action:   exp1155Token_SafeTransferFromBatch,
		},
		{
			Name:     "burn",
			Usage:    "<from> <id> <amount>",
			Category: "arithmetic",
			Action:   exp1155Token_Burn,
		},
		{
			Name:     "burnBatch",
			Usage:    "<from> <ids> <amounts>",
			Category: "arithmetic",
			Action:   exp1155Token_BurnBatch,
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
	amount, ok := big.NewInt(0).SetString(args.Get(2), 0)
	if !ok {
		fmt.Println(c.App.Usage)
		return nil
	}
	Exp1155Class := &reqcontract.Exp1155MintArgs{
		Recipient: args.Get(0),
		Tokenurl:  args.Get(1),
		Amount:    amount,
	}
	txhash, err := defaultEXP1155Token.Mint(Exp1155Class)
	if err != nil {
		return err
	}
	fmt.Println(txhash)
	return nil
}

func exp1155Token_MintBatch(c *cli.Context) error {
	args := c.Args()

	amounts := strings.Split(args.Get(1), ",")
	tokenurls := strings.Split(args.Get(2), ",")
	if len(amounts) != len(tokenurls) {
		fmt.Println(c.App.Usage)
		return nil
	}
	exp1155Class := &reqcontract.Exp1155MintBatchArgs{
		Recipient: args.Get(0),
		Amounts:   args.Get(1),
		TokenUrls: args.Get(2),
	}

	txhash, err := defaultEXP1155Token.MintBatch(exp1155Class)
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
	result, err := defaultEXP1155Token.BalanceOf(args.Get(0), tokenid)
	if err != nil {
		return err
	}
	fmt.Printf("result:%v\n", result)
	return nil
}

func exp1155Token_BalanceOfBatch(c *cli.Context) error {
	args := c.Args()

	addrs := strings.Split(args.Get(0), ",")
	tokenids := strings.Split(args.Get(1), ",")
	result, err := defaultEXP1155Token.BalanceOfBatch(addrs, tokenids)
	if err != nil {
		return err
	}
	fmt.Printf("result:%v\n", result)
	return nil
}

func exp1155Token_SetApprovalForAll(c *cli.Context) error {
	args := c.Args()
	if c.NArg() < 3 {
		fmt.Println(c.App.Usage)
		return nil
	}
	approved, err := strconv.ParseBool(args.Get(1))
	if err != nil {
		return err
	}
	req := &reqcontract.Exp1155SetApprovalForAllArgs{
		Operator: args.Get(0),
		Approved: approved,
		PriKey:   args.Get(2),
	}
	txhash, err := defaultEXP1155Token.SetApprovalForAll(req)
	if err != nil {
		return err
	}
	fmt.Printf("txhash:%v\n", txhash)
	return nil
}

func exp1155Token_IsApprovalForAll(c *cli.Context) error {
	args := c.Args()
	if c.NArg() < 2 {
		fmt.Println(c.App.Usage)
		return nil
	}
	req := &reqcontract.Exp1155IsApprovedForAllArgs{
		Account:  args.Get(0),
		Operator: args.Get(1),
	}
	approved, err := defaultEXP1155Token.IsApprovalForAll(req)
	if err != nil {
		return err
	}
	fmt.Printf("%v\n", approved)
	return nil
}

func exp1155Token_SafeTransferFrom(c *cli.Context) error {
	args := c.Args()

	if c.NArg() < 4 {
		fmt.Println(c.App.Usage)
		return nil
	}
	tokenid, ok := big.NewInt(0).SetString(args.Get(2), 0)
	if !ok {
		fmt.Println(c.App.Usage)
		return nil
	}
	amount, ok := big.NewInt(0).SetString(args.Get(3), 0)
	if !ok {
		fmt.Println(c.App.Usage)
		return nil
	}
	req := &reqcontract.Exp1155SafeTransferFromArgs{
		From:   args.Get(0),
		To:     args.Get(1),
		Id:     tokenid,
		Amount: amount,
		Prikey: args.Get(4),
	}
	txhash, err := defaultEXP1155Token.TransferFrom(req)
	if err != nil {
		return err
	}
	fmt.Printf("txhash:%v\n", txhash)
	return nil
}

func exp1155Token_SafeTransferFromBatch(c *cli.Context) error {
	args := c.Args()
	if c.NArg() < 4 {
		fmt.Println(c.App.Usage)
		return nil
	}

	tokenids := strings.Split(args.Get(2), ",")
	amounts := strings.Split(args.Get(3), ",")
	if len(tokenids) == len(amounts) {
		fmt.Println(c.App.Usage)
		return nil
	}
	req := &reqcontract.Exp1155SafeBatchTransferFromArgs{
		From:    args.Get(0),
		To:      args.Get(1),
		Ids:     args.Get(2),
		Amounts: args.Get(3),
		Prikey:  args.Get(4),
	}
	txhash, err := defaultEXP1155Token.TransferFromBatch(req)
	if err != nil {
		return err
	}
	fmt.Printf("txhash:%v\n", txhash)
	return nil
}

func exp1155Token_Burn(c *cli.Context) error {
	args := c.Args()
	if c.NArg() < 3 {
		fmt.Println(c.App.Usage)
		return nil
	}

	tokenid, ok := big.NewInt(0).SetString(args.Get(1), 0)
	if !ok {
		fmt.Println(c.App.Usage)
		return nil
	}
	amount, ok := big.NewInt(0).SetString(args.Get(2), 0)
	if !ok {
		fmt.Println(c.App.Usage)
		return nil
	}
	req := &reqcontract.Exp1155BurnArgs{
		From:   args.Get(0),
		ID:     tokenid,
		Amount: amount,
		Prikey: args.Get(3),
	}
	txhash, err := defaultEXP1155Token.Burn(req)
	if err != nil {
		return err
	}
	fmt.Printf("txhash:%v\n", txhash)
	return nil
}

func exp1155Token_BurnBatch(c *cli.Context) error {
	args := c.Args()
	if c.NArg() < 3 {
		fmt.Println(c.App.Usage)
		return nil
	}

	tokenids := strings.Split(args.Get(1), ",")
	amounts := strings.Split(args.Get(2), ",")
	if len(tokenids) != len(amounts) {
		fmt.Println(c.App.Usage)
		return nil
	}
	req := &reqcontract.Exp1155BurnBatchArgs{
		From:    args.Get(0),
		IDs:     args.Get(1),
		Amounts: args.Get(2),
		Prikey:  args.Get(3),
	}
	txhash, err := defaultEXP1155Token.BurnBatch(req)
	if err != nil {
		return err
	}
	fmt.Printf("txhash:%v\n", txhash)
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
