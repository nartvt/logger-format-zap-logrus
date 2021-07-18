package main

import (
	"logrus-format/logger"
)

func init() {
	logger.InitEasyFormat()
}

func main() {
	zapLog := logger.ZapFormat()
	logger.LoggerZap = logger.NewZapLogger(zapLog)
	testLog()
	defer zapLog.Sync()
	//logger.Infof("Message Testing")
}

func testLog() {
	logZap := logger.LoggerZap.GetZapLogger()
	logZap.Info("Message Testting")
}
