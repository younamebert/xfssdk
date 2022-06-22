package main

import (
	"fmt"
	"log"
	"os"

	"github.com/younamebert/xfssdk"
	"github.com/younamebert/xfssdk/contract/stdtoken"
	reqcontract "github.com/younamebert/xfssdk/servce/contract/request"
	"gopkg.in/urfave/cli.v1"
)

var (
	Key           = "0x0101e0f42d6125c515fa1065875c665d22ac23cb2c9457d36e9dc4487cd873bad1c3"
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
			// After: Stdtoken_Create,
		},
		{
			Name: "deploy",
			// Aliases:  []string{"create"},
			Usage:    "deploy <code> <addrprikey>",
			Category: "arithmetic",
			Action:   Stdtoken_Deploy,
			// After: Stdtoken_Create,
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
