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
// sugar zap
func ConfigZap() *zap.SugaredLogger{

	cfg := zap.Config{
		Encoding:    "json", //encode kiểu json hoặc console
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),   //chọn InfoLevel có thể log ở cả 3 level
		OutputPaths: []string{"stderr"},

		EncoderConfig: zapcore.EncoderConfig{   //Cấu hình logging, sẽ không có stacktracekey
			MessageKey: "message",
			TimeKey: "time",
			LevelKey: "level",
			CallerKey:    "caller",
			EncodeCaller: zapcore.FullCallerEncoder,    //Lấy dòng code bắt đầu log
			EncodeLevel: CustomLevelEncoder,    //Format cách hiển thị level log
			EncodeTime: SyslogTimeEncoder,  //Format hiển thị thời điểm log
		},
	}

	logger, _ := cfg.Build()    //Build ra Logger
	return logger.Sugar()   //Trả về logger hoặc Sugaredlogger, ở đây ta chọn trả về Logger
}