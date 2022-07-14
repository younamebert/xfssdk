package main

import (
	"fmt"
	"log"
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
	BridgeKey           = "0x010170c10441b71cfdf7af3ed299151be8668c9c79f99ec68f8e34b2ee1151c46a80"
	DefaultBridgeAddr   = crypto.Prikey2Addr(BridgeKey)
	deafaultBridgetoken = &bridge.Bridge{
		Bankaddress:          "YyhP7nipqxj8eoUN6Lj1P4q8vdQ8mwgVX",
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
			Usage:    "create <name> <symbol> <caddrss>",
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
			Usage:    "<to> <amount>",
			Category: "arithmetic",
			Action:   Bridge_TransferIn,
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
	argsCreate := &reqcontract.BridgeArgs{
		Name:            args.Get(0),
		Symbol:          args.Get(1),
		ContractAddress: args.Get(2),
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
		TransferFromAddressPriKey: BridgeKey,
		TransferToAddress:         args.Get(0),
		TransferAmount:            args.Get(1),
	}
	txhash, err := deafaultBridgetoken.TransferIn(argsTransfer)
	if err != nil {
		return err
	}
	fmt.Printf("txhahs:%v\n", txhash)
	return nil
}
