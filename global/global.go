package global

import (
	"go.uber.org/zap"
)

var GVA_LOG *zap.Logger

func Set_GVA_LOG(logger *zap.Logger) {
	GVA_LOG = logger
}
