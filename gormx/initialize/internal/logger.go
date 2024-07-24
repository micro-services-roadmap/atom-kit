package internal

import (
	"fmt"
	"github.com/micro-services-roadmap/kit-common/kg"
	"gorm.io/gorm/logger"
)

const (
	Console = "console"
	Zap     = "zap"
	GoZero  = "go-zero"
)

type writer struct {
	logger.Writer
}

// NewWriter writer 构造函数
func NewWriter(w logger.Writer) *writer {
	return &writer{Writer: w}
}

// Printf 格式化打印日志
func (w *writer) Printf(message string, data ...interface{}) {
	var logType string
	switch kg.C.System.DbType {
	case "mysql":
		logType = kg.C.Mysql.LogType
	case "pgsql":
		logType = kg.C.Pgsql.LogType
	}

	if logType == Zap {
		kg.L.Info(fmt.Sprintf(message+"\n", data...))
	} else if logType == GoZero {
		// logx.Debug(fmt.Sprintf(message+"\n", data...))
	}

	w.Writer.Printf(message, data...)
}
