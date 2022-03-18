package xfssdk

import (
	"github.com/younamebert/xfssdk/api"
	"github.com/younamebert/xfssdk/config"
	"github.com/younamebert/xfssdk/core"
	"github.com/younamebert/xfssdk/core/apis"
	"github.com/younamebert/xfssdk/exactly"
	"github.com/younamebert/xfssdk/global"
	"github.com/younamebert/xfssdk/libs/client"
)

type Handle struct {
	Config    *config.HandleConfig
	Exactly   *exactly.Exactly
	ApiMethod *api.ApiMethod
}

// Default  create default handle object
func Default() *Handle {
	handle := &Handle{
		Config:    config.DefaultHandleConfig(),
		Exactly:   exactly.NewExactly(),
		ApiMethod: api.NewApiMethod(),
	}
	cli := client.NewClient(handle.Config.NodeLink, handle.Config.NodeLinkOutTime)
	apis.SetXFSClient(cli)

	XFSLogger := core.NewXFSLogger(config.DefaultLoggerConfig())
	logc := XFSLogger.Zap()
	global.Set_GVA_LOG(logc)

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

	handle.ApiMethod = api.NewApiMethod()
	handle.Exactly = exactly.NewExactly()
	return handle
}
