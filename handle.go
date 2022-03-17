package xfssdk

import (
	"xfssdk/api"
	"xfssdk/config"
	"xfssdk/core"
	"xfssdk/core/apis"
	"xfssdk/exactly"
	"xfssdk/global"
	"xfssdk/libs/client"
)

type Handle struct {
	Config    *config.HandleConfig
	Exactly   *exactly.Exactly
	ApiMethod *api.ApiMethod
}

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
