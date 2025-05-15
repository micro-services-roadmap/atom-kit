package internal

import (
	"github.com/micro-services-roadmap/atom-kit/zapx/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Zap = new(_zap)

type _zap struct{}

// GetEncoder 获取 zapx-core.Encoder
func (z *_zap) GetEncoder(config config.Zap) zapcore.Encoder {
	if config.Format == "json" {
		return zapcore.NewJSONEncoder(z.GetEncoderConfig(config))
	}
	return zapcore.NewConsoleEncoder(z.GetEncoderConfig(config))
}

// GetEncoderConfig 获取zapcore.EncoderConfig
func (z *_zap) GetEncoderConfig(config config.Zap) zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  config.StacktraceKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    config.ZapEncodeLevel(),
		EncodeTime:     config.CustomTimeEncoder(),
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}
}

// GetEncoderCore 获取Encoder的 zapcore.Core
func (z *_zap) GetEncoderCore(config config.Zap, l zapcore.Level, level zap.LevelEnablerFunc) zapcore.Core {
	writer := FileRotatelogs.GetWriteSyncer(config, l.String()) // 日志分割
	return zapcore.NewCore(z.GetEncoder(config), writer, level)
}

// GetZapCores 根据配置文件的Level获取 []zapcore.Core
func (z *_zap) GetZapCores(config config.Zap) []zapcore.Core {
	cores := make([]zapcore.Core, 0, 7)
	for level := config.TransportLevel(); level <= zapcore.FatalLevel; level++ {
		cores = append(cores, z.GetEncoderCore(config, level, z.GetLevelPriority(level)))
	}

	return cores
}

// GetLevelPriority 根据 zapcore.Level 获取 zap.LevelEnablerFunc
func (z *_zap) GetLevelPriority(level zapcore.Level) zap.LevelEnablerFunc {
	switch level {
	case zapcore.DebugLevel:
		return func(level zapcore.Level) bool { // 调试级别
			return level == zap.DebugLevel
		}
	case zapcore.InfoLevel:
		return func(level zapcore.Level) bool { // 日志级别
			return level == zap.InfoLevel
		}
	case zapcore.WarnLevel:
		return func(level zapcore.Level) bool { // 警告级别
			return level == zap.WarnLevel
		}
	case zapcore.ErrorLevel:
		return func(level zapcore.Level) bool { // 错误级别
			return level == zap.ErrorLevel
		}
	case zapcore.DPanicLevel:
		return func(level zapcore.Level) bool { // dpanic级别
			return level == zap.DPanicLevel
		}
	case zapcore.PanicLevel:
		return func(level zapcore.Level) bool { // panic级别
			return level == zap.PanicLevel
		}
	case zapcore.FatalLevel:
		return func(level zapcore.Level) bool { // 终止级别
			return level == zap.FatalLevel
		}
	default:
		return func(level zapcore.Level) bool { // 调试级别
			return level == zap.DebugLevel
		}
	}
}
