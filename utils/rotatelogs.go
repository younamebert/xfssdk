package utils

import (
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap/zapcore"
)

func GetWriteSyncer(LogInConsole bool, file string) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   file, // Location of log files
		MaxSize:    10,   // The maximum size (in megabytes) of the log file before cutting
		MaxBackups: 200,  // Maximum number of old files retained
		MaxAge:     30,   // Maximum number of days to keep old files
		Compress:   true, // Compress / archive old files
	}

	if LogInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(lumberJackLogger))
	}
	return zapcore.AddSync(lumberJackLogger)
}
