package main

import (
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"

	"github.com/younamebert/xfssdk"
	"github.com/younamebert/xfssdk/common"
	"github.com/younamebert/xfssdk/contract/bridge"
	"github.com/younamebert/xfssdk/crypto"
	reqcontract "github.com/younamebert/xfssdk/servce/contract/request"
	"gopkg.in/urfave/cli.v1"
)

var (
	BridgeKey           = "0x01010166b084b55d17e8619635cf9131e29bf732f9a05c2c8cc2a11ed9f8e3d4cf1b"
	DefaultBridgeAddr   = crypto.Prikey2Addr(BridgeKey)
	deafaultBridgetoken = &bridge.Bridge{
		Bankaddress:          "WfSHAkLLS7rn1n6Cev1Q5wj3wxMBAGJ66",
		CreatorAddressPrikey: BridgeKey,
	}
	app         *cli.App
	handle      = xfssdk.Default()
	bridgelocal = new(bridge.BridgeLocal)
)

// var app = cli.NewApp()

func init() {
	app = cli.NewApp()
	app.Name = "bridge"
	app.Usage = "bridge usdt/xfs"
	app.Version = "1.0.0"
	// handle.ContractEngine.StdToken = new(stdtoken.StdTokenLocal)
}

func main() {
	app.Commands = []cli.Command{
		{
			Name:     "create",
			Usage:    "create <name> <symbol> <caddrss> <chainid>",
			Category: "arithmetic",
			Action:   Bridge_Create,
		},
		{
			Name:     "deploy",
			Usage:    "deploy <code> <addrprikey>",
			Category: "arithmetic",
			Action:   Bridge_Deploy,
		},
		{
			Name:     "caddr",
			Usage:    "<address> <nonce>",
			Category: "arithmetic",
			Action:   Bridge_caddr,
		},
		{
			Name:     "transferIn",
			Usage:    "<prikey> <to> <amount> <fromChainId>",
			Category: "arithmetic",
			Action:   Bridge_TransferIn,
		},
		{
			Name:     "transferOut",
			Usage:    "<prikey> <address> <toaddress> <amount> <toChainId>",
			Category: "arithmetic",
			Action:   Bridge_TransferOut,
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func Bridge_Create(c *cli.Context) error {
	args := c.Args()

	if c.NArg() < 3 {
		fmt.Println(c.App.Usage)
		return nil
	}
	chainid, _ := new(big.Int).SetString(args.Get(3), 10)
	argsCreate := &reqcontract.BridgeArgs{
		Name:            args.Get(0),
		Symbol:          args.Get(1),
		ContractAddress: args.Get(2),
		ChainId:         chainid,
	}
	code, err := bridgelocal.Create(argsCreate)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return err
	}
	fmt.Println(code)
	return nil
}

func Bridge_Deploy(c *cli.Context) error {
	args := c.Args()

	if c.NArg() < 1 {
		fmt.Println(c.App.Usage)
		return nil
	}

	argsDeploy := reqcontract.DeployBridgeArgs{
		Code:       args.Get(0),
		Addresskey: BridgeKey,
	}
	_, txhash, err := bridgelocal.DeployDridge(argsDeploy)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return err
	}
	fmt.Println(txhash)
	return nil
}

func Bridge_caddr(c *cli.Context) error {
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

func Bridge_TransferIn(c *cli.Context) error {
	args := c.Args()
	argsTransfer := reqcontract.BridgeTransferInArgs{
		TransferFromAddressPriKey: args.Get(0),
		DepositorAddress:          args.Get(1),
		TransferAmount:            args.Get(2),
		TransferFromChainId:       args.Get(3),
	}
	txhash, err := deafaultBridgetoken.TransferIn(argsTransfer)
	if err != nil {
		return err
	}
	fmt.Printf("txhash:%v\n", txhash)
	return nil
}

func Bridge_TransferOut(c *cli.Context) error {
	args := c.Args()
	argsTransfer := reqcontract.BridgeTransferOutArgs{
		TransferFromAddressPriKey: args.Get(0),
		TransferAddress:           args.Get(1),
		TransferToAddress:         args.Get(2),
		TransferAmount:            args.Get(3),
		TransferToChainId:         args.Get(4),
	}
	txhash, err := deafaultBridgetoken.TransferOut(argsTransfer)
	if err != nil {
		return err
	}
	fmt.Printf("txhash:%v\n", txhash)
	return nil
}
