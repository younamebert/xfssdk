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
	Key           = "0x01010e2f7fdf0c76d4dfaeb76a30aa1a98dcaf5d10dc9596cbafaf5806c14074a813"
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
