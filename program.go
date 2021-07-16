package main

import (
	"logrus-format/logger"
)

func init() {
	logger.InitEasyFormat()
}
func main() {
	logger.Infof("Message test line and func")
}
