package config

import (
	"strings"
	"time"
)

type HandleConfig struct {
	Networkid       uint32 `json:"networkid" yaml:"networkid"`             //explain:networkid default:1
	Version         uint32 `json:"version" yaml:"version"`                 //explain:version default:3
	NodeLink        string `json:"nodelink" yaml:"nodelink"`               //explain:connection request address default:https://api.scan.xfs.tech/jsonrpc/v2
	NodeLinkOutTime string `json:"nodelinkouttime" yaml:"nodelinkouttime"` //explain:timeout for connection request clients default:5s
	Logger          *LoggerConfig
}

func NewHandleConfig(handleconfig *HandleConfig) *HandleConfig {
	config := DefaultHandleConfig()

	if handleconfig.Networkid != 0 {
		config.SetNetworkid(handleconfig.Networkid)
	}

	if handleconfig.NodeLink != "" {
		if strings.Contains(handleconfig.NodeLink, "http://") || strings.Contains(handleconfig.NodeLink, "https://") {
			config.SetNodeLink(handleconfig.NodeLink)
		}
	}

	if handleconfig.NodeLinkOutTime != "" {
		_, err := time.ParseDuration(handleconfig.NodeLinkOutTime)
		if err == nil {
			config.SetNodeLinkOutTime(handleconfig.NodeLinkOutTime)
		}
	}
	return config
}

func DefaultHandleConfig() *HandleConfig {
	return &HandleConfig{
		Networkid:       1,
		Version:         1,
		NodeLink:        "http://127.0.0.1:9012/", // http://192.168.2.13:9014/
		NodeLinkOutTime: "5s",
		Logger:          DefaultLoggerConfig(),
	}
}

func (handle *HandleConfig) SetNetworkid(networkid uint32) {
	handle.Networkid = networkid
}

func (handle *HandleConfig) SetNodeLink(nodelink string) {
	handle.NodeLink = nodelink
}

func (handle *HandleConfig) SetNodeLinkOutTime(outtime string) {
	handle.NodeLinkOutTime = outtime
}

func (handle *HandleConfig) SetVersion(version uint32) {
	handle.Version = version
}

type LoggerConfig struct {
	Level         string `json:"level" yaml:"level"`                  //explain:Log level
	Format        string `json:"format" yaml:"format"`                //explain:Log printing
	Prefix        string `json:"prefix" yaml:"prefix"`                //explain:Log prefix
	Director      string `json:"director"  yaml:"director"`           //explain:Log folder
	ShowLine      bool   `json:"showLine" yaml:"showLine"`            //explain:Display line
	EncodeLevel   string `json:"encodeLevel" yaml:"encode-level"`     //explain:Coding level
	StacktraceKey string `json:"stacktraceKey" yaml:"stacktrace-key"` //explain:Stack name
	LogInConsole  bool   `json:"logInConsole" yaml:"log-in-console"`  //explain:Print to console
}

const (
	LowercaseLevelEncoderCode      = "LowercaseLevelEncoder"
	LowercaseColorLevelEncoderCode = "LowercaseColorLevelEncoder"
	CapitalLevelEncoderCode        = "CapitalLevelEncoder"
	CapitalColorLevelEncoderCode   = "CapitalColorLevelEncoder"
)

func DefaultLoggerConfig() *LoggerConfig {
	return &LoggerConfig{
		Level:         "info",
		Format:        "console",
		Prefix:        "xfssdk",
		Director:      "log",
		ShowLine:      true,
		EncodeLevel:   "LowercaseColorLevelEncoder",
		StacktraceKey: "stacktrace",
		LogInConsole:  true,
	}
}
