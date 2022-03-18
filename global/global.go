package global

import (
	"go.uber.org/zap"
)

var GVA_LOG *zap.Logger

// Set_GVA_LOG 设置全局日志
func Set_GVA_LOG(logger *zap.Logger) {
	GVA_LOG = logger
}
