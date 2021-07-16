package logger

import (
	log "github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
	prefixed "github.com/t-tomalak/logrus-prefixed-formatter"
)

func initLogFormat() {
	formatter := &prefixed.TextFormatter{
		DisableColors:   true,
		ForceColors:     true,
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
		ForceFormatting: true,
	}
	log.SetFormatter(formatter)
}

func initEasy() {
	formatter := &easy.Formatter{

		TimestampFormat: "2006-01-02 15:04:05",
		LogFormat:       "%time% [%lvl%] - %msg% ",
	}
	log.SetFormatter(formatter)
}
