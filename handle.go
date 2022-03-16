package xfssdk

import (
	"xfssdk/api"
	"xfssdk/core/apis"
	"xfssdk/exactly"
	"xfssdk/libs/client"
)

type Handle struct {
	Config    *HandleConfig
	Exactly   *exactly.Exactly
	ApiMethod *api.ApiMethod
}

func Default() *Handle {

	handle := &Handle{
		Config:    DefaultHandleConfig(),
		Exactly:   exactly.NewExactly(),
		ApiMethod: api.NewApiMethod(),
	}
	cli := client.NewClient(handle.Config.NodeLink, handle.Config.NodeLinkOutTime)
	apis.SetXFSClient(cli)
	return handle
}

func New(handleconf *HandleConfig, loggerconf *LoggerConfig) *Handle {
	handle := new(Handle)
	handle.Config = NewHandleConfig(handleconf, loggerconf)
	cli := client.NewClient(handle.Config.NodeLink, handle.Config.NodeLinkOutTime)
	apis.SetXFSClient(cli)

	handle.ApiMethod = api.NewApiMethod()
	handle.Exactly = exactly.NewExactly()
	return handle
}
