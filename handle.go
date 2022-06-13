package xfssdk

import (
	"fmt"
	"os"

	"github.com/younamebert/xfssdk/api"
	"github.com/younamebert/xfssdk/config"
	"github.com/younamebert/xfssdk/contract"
	"github.com/younamebert/xfssdk/core"
	"github.com/younamebert/xfssdk/core/apis"
	"github.com/younamebert/xfssdk/global"
	"github.com/younamebert/xfssdk/libs/client"
)

type Handle struct {
	Config         *config.HandleConfig
	ApiMethod      *api.ApiMethod
	ContractEngine *contract.ContractEngine
}

// Default  create default handle object
func Default() *Handle {
	handle := &Handle{
		Config:    config.DefaultHandleConfig(),
		ApiMethod: api.NewApiMethod(),
	}
	cli := client.NewClient(handle.Config.NodeLink, handle.Config.NodeLinkOutTime)
	apis.SetXFSClient(cli)

	XFSLogger := core.NewXFSLogger(config.DefaultLoggerConfig())
	logc := XFSLogger.Zap()
	global.Set_GVA_LOG(logc)
	// Initialize the ABI object
	if err := apis.XFSABI(); err != nil {
		fmt.Println(err)
		os.Exit(1)
		return nil
	}
	return handle
}

// New adopt handleconfig and loggerconfig create handle object
func New(handleconf *config.HandleConfig, loggerconf *config.LoggerConfig) *Handle {
	handle := new(Handle)
	handle.Config = config.NewHandleConfig(handleconf, loggerconf)
	cli := client.NewClient(handle.Config.NodeLink, handle.Config.NodeLinkOutTime)
	apis.SetXFSClient(cli)

	XFSLogger := core.NewXFSLogger(handle.Config.Logger)
	logc := XFSLogger.Zap()
	global.Set_GVA_LOG(logc)
	// Initialize the ABI object
	if err := apis.XFSABI(); err != nil {
		fmt.Println(err)
		os.Exit(1)
		return nil
	}

	handle.ApiMethod = api.NewApiMethod()
	return handle
}
