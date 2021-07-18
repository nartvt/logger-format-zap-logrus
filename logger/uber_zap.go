package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

type zapLoggerCustom struct {
	loggerZap *zap.Logger
}

var (
	LoggerZap *zapLoggerCustom
)

func NewZapLogger(logger *zap.Logger) *zapLoggerCustom {
	log := &zapLoggerCustom{}
	if logger != nil {
		log.loggerZap = logger
	}
	return log
}

func (zap *zapLoggerCustom) GetZapLogger() *zap.Logger {
	return zap.loggerZap
}

func ZapFormat() *zap.Logger {
	cfg := &zap.Config{
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		Encoding:    "console",
		OutputPaths: []string{"stderr"},
		//OutputPaths: []string{"stdout", "stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:   "message",
			TimeKey:      "time",
			LevelKey:     "level",
			CallerKey:    "caller",
			EncodeCaller: zapcore.FullCallerEncoder,
			EncodeTime:   SyslogTimeEncoder,
			EncodeLevel:  CustomLevelEncoder,
		},
	}
	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	// defer logger.Sync()
	return logger
}

func SyslogTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}

func CustomLevelEncoder(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + level.CapitalString() + "]")
}
