package logger

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
	prefixed "github.com/t-tomalak/logrus-prefixed-formatter"
	"runtime"
	"sort"
	"strings"
)

func InitEasyFormat() {
	formatter := &prefixed.TextFormatter{
		DisableColors:   true,
		ForceColors:     true,
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
		ForceFormatting: true,
	}
	log.SetFormatter(formatter)
}

func InitLogFormat() {
	log.SetReportCaller(true)
	//logger := log.New()
	//buffer := &bytes.Buffer{}
	//logger.Out = buffer
	//logger.SetReportCaller(true)
	textFormat := &log.TextFormatter{
		//DisableColors: true,
		TimestampFormat:  "2006-01-02 15:04:05",
		CallerPrettyfier: runTime,
		FullTimestamp:    true,
		SortingFunc: func(s []string) {
			sort.Strings(s)
		},
		DisableSorting: false,
		ForceColors:    false,
		ForceQuote:     false,
	}
	//entry := &log.Entry{
	//	Logger: logger,
	//}
	//textFormat.Format(entry)
	log.SetFormatter(textFormat)
}
func InitEasy() {
	formatter := &easy.Formatter{

		TimestampFormat: "2006-01-02 15:04:05",
		LogFormat:       "%time% [%lvl%] - %msg% ",
	}
	log.SetFormatter(formatter)
}

func fileLine(toFile string) string {
	arr := strings.Split(toFile, "/")
	return arr[len(arr)-1]
}

func runTime(f *runtime.Frame) (string, string) {
	return "", fmt.Sprintf("\t%s - line %d", fileLine(f.File), f.Line)
}

