package xfssdk

import (
	"fmt"
	"os"

	"github.com/younamebert/xfssdk/api"
	"github.com/younamebert/xfssdk/config"
	"github.com/younamebert/xfssdk/core/apis"
	"github.com/younamebert/xfssdk/libs/client"
)

type Handle struct {
	Config    *config.HandleConfig
	ApiMethod *api.ApiMethod
}

// Default  create default handle object
func Default() *Handle {
	handle := &Handle{
		Config:    config.DefaultHandleConfig(),
		ApiMethod: api.NewApiMethod(),
	}
	cli := client.NewClient(handle.Config.NodeLink, handle.Config.NodeLinkOutTime)
	apis.SetXFSClient(cli)
	// Initialize the ABI object
	if err := apis.XFSABI(); err != nil {
		fmt.Println(err)
		os.Exit(1)
		return nil
	}
	return handle
}

// New adopt handleconfig and loggerconfig create handle object
func New(handleconf *config.HandleConfig) *Handle {
	handle := new(Handle)
	handle.Config = config.NewHandleConfig(handleconf)
	cli := client.NewClient(handle.Config.NodeLink, handle.Config.NodeLinkOutTime)
	apis.SetXFSClient(cli)

	// Initialize the ABI object
	if err := apis.XFSABI(); err != nil {
		fmt.Println(err)
		os.Exit(1)
		return nil
	}

	handle.ApiMethod = api.NewApiMethod()
	return handle
}

func SetupGLobal(link, reqLinkOutTime string) error {
	cli := client.NewClient(link, reqLinkOutTime)
	apis.SetXFSClient(cli)
	if err := apis.XFSABI(); err != nil {
		fmt.Println(err)
		os.Exit(1)
		return nil
	}
	return nil
}
