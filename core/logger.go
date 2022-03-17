package core

import (
	"errors"
	"fmt"
	"os"
	"time"
	"xfssdk/config"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type XFSLogger struct {
	conf *config.LoggerConfig
}

func NewXFSLogger(conf *config.LoggerConfig) *XFSLogger {
	return &XFSLogger{
		conf: conf,
	}
}

func (xfslog *XFSLogger) Zap() (logger *zap.Logger) {
	if ok, _ := PathExists(xfslog.conf.Director); !ok { // 判断是否有Director文件夹
		_ = os.Mkdir(xfslog.conf.Director, os.ModePerm)
	}
	// 调试级别
	debugPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.DebugLevel
	})
	// 日志级别
	infoPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.InfoLevel
	})
	// 警告级别
	warnPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.WarnLevel
	})
	// 错误级别
	errorPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev >= zap.ErrorLevel
	})

	cores := [...]zapcore.Core{
		xfslog.getEncoderCore(fmt.Sprintf("./%s/server_debug.log", xfslog.conf.Director), debugPriority),
		xfslog.getEncoderCore(fmt.Sprintf("./%s/server_info.log", xfslog.conf.Director), infoPriority),
		xfslog.getEncoderCore(fmt.Sprintf("./%s/server_warn.log", xfslog.conf.Director), warnPriority),
		xfslog.getEncoderCore(fmt.Sprintf("./%s/server_error.log", xfslog.conf.Director), errorPriority),
	}
	logger = zap.New(zapcore.NewTee(cores[:]...), zap.AddCaller())

	if xfslog.conf.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}

// getEncoderConfig 获取zapcore.EncoderConfig
func (xfslog *XFSLogger) getEncoderConfig() (config zapcore.EncoderConfig) {
	config = zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  xfslog.conf.StacktraceKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     xfslog.CustomTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}
	switch {
	case xfslog.conf.EncodeLevel == "LowercaseLevelEncoder": // 小写编码器(默认)
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	case xfslog.conf.EncodeLevel == "LowercaseColorLevelEncoder": // 小写编码器带颜色
		config.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	case xfslog.conf.EncodeLevel == "CapitalLevelEncoder": // 大写编码器
		config.EncodeLevel = zapcore.CapitalLevelEncoder
	case xfslog.conf.EncodeLevel == "CapitalColorLevelEncoder": // 大写编码器带颜色
		config.EncodeLevel = zapcore.CapitalColorLevelEncoder
	default:
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	}
	return config
}

// getEncoder 获取zapcore.Encoder
func (xfslog *XFSLogger) getEncoder() zapcore.Encoder {
	if xfslog.conf.Format == "json" {
		return zapcore.NewJSONEncoder(xfslog.getEncoderConfig())
	}
	return zapcore.NewConsoleEncoder(xfslog.getEncoderConfig())
}

// getEncoderCore 获取Encoder的zapcore.Core
func (xfslog *XFSLogger) getEncoderCore(fileName string, level zapcore.LevelEnabler) (core zapcore.Core) {
	writer := GetWriteSyncer(xfslog.conf.LogInConsole, fileName) // 使用file-rotatelogs进行日志分割
	return zapcore.NewCore(xfslog.getEncoder(), writer, level)
}

// 自定义日志输出时间格式
func (xfslog *XFSLogger) CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(xfslog.conf.Prefix + "2006/01/02 - 15:04:05.000"))
}

func PathExists(path string) (bool, error) {
	fi, err := os.Stat(path)
	if err == nil {
		if fi.IsDir() {
			return true, nil
		}
		return false, errors.New("存在同名文件")
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func CreateDir(dirs ...string) (err error) {
	for _, v := range dirs {
		exist, err := PathExists(v)
		if err != nil {
			return err
		}
		if !exist {
			// logger.GVA_LOG.Debug("create directory" + v)
			if err := os.MkdirAll(v, os.ModePerm); err != nil {
				// logger.GVA_LOG.Error("create directory"+v, zap.Any(" error:", err))
				return err
			}
		}
	}
	return err
}

func GetWriteSyncer(LogInConsole bool, file string) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   file, // 日志文件的位置
		MaxSize:    10,   // 在进行切割之前，日志文件的最大大小（以MB为单位）
		MaxBackups: 200,  // 保留旧文件的最大个数
		MaxAge:     30,   // 保留旧文件的最大天数
		Compress:   true, // 是否压缩/归档旧文件
	}

	if LogInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(lumberJackLogger))
	}
	return zapcore.AddSync(lumberJackLogger)
}
