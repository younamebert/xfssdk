package core

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/younamebert/xfssdk/config"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type XFSLogger struct {
	conf *config.LoggerConfig
}

// NewXFSLogger Initialize
func NewXFSLogger(conf *config.LoggerConfig) *XFSLogger {
	return &XFSLogger{
		conf: conf,
	}
}

func (xfslog *XFSLogger) Zap() (logger *zap.Logger) {
	if ok, _ := PathExists(xfslog.conf.Director); !ok { // Determine whether there is a director folder
		_ = os.Mkdir(xfslog.conf.Director, os.ModePerm)
	}
	// debug level
	debugPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.DebugLevel
	})
	// log level
	infoPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.InfoLevel
	})
	// warning level
	warnPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.WarnLevel
	})
	// error level
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

// getEncoderConfig get zapcore.EncoderConfig
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
	case xfslog.conf.EncodeLevel == "LowercaseLevelEncoder": //Lower case encoder (default)
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	case xfslog.conf.EncodeLevel == "LowercaseColorLevelEncoder": //Lowercase encoder with color
		config.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	case xfslog.conf.EncodeLevel == "CapitalLevelEncoder": //Uppercase encoder
		config.EncodeLevel = zapcore.CapitalLevelEncoder
	case xfslog.conf.EncodeLevel == "CapitalColorLevelEncoder": //Uppercase encoder with color
		config.EncodeLevel = zapcore.CapitalColorLevelEncoder
	default:
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	}
	return config
}

// getEncoder get zapcore.Encoder
func (xfslog *XFSLogger) getEncoder() zapcore.Encoder {
	if xfslog.conf.Format == "json" {
		return zapcore.NewJSONEncoder(xfslog.getEncoderConfig())
	}
	return zapcore.NewConsoleEncoder(xfslog.getEncoderConfig())
}

// getEncoderCore get Encoder->zapcore.Core
func (xfslog *XFSLogger) getEncoderCore(fileName string, level zapcore.LevelEnabler) (core zapcore.Core) {
	writer := GetWriteSyncer(xfslog.conf.LogInConsole, fileName) //Log segmentation using file rotatelogs
	return zapcore.NewCore(xfslog.getEncoder(), writer, level)
}

//custom log output time format
func (xfslog *XFSLogger) CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(xfslog.conf.Prefix + "2006/01/02 - 15:04:05.000"))
}

func PathExists(path string) (bool, error) {
	fi, err := os.Stat(path)
	if err == nil {
		if fi.IsDir() {
			return true, nil
		}
		return false, errors.New("a file with the same name already exists")
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
		Filename:   file, // Location of log files
		MaxSize:    10,   // The maximum size (in MB) of the log file before cutting
		MaxBackups: 200,  // Maximum number of old files retained
		MaxAge:     30,   // Maximum number of days to keep old files
		Compress:   true, // Compress / archive old files
	}

	if LogInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(lumberJackLogger))
	}
	return zapcore.AddSync(lumberJackLogger)
}
