package gen

import (
	"gorm-gen/mylog"

	"go.uber.org/zap"
)

var Logger *zap.SugaredLogger = mylog.Logger
