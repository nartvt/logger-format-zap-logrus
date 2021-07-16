package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
	prefixed "github.com/t-tomalak/logrus-prefixed-formatter"
	"regexp"
	"runtime"
	"sort"
	"strings"
)

func initLogFormat() {
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
func initEasy() {
	//pc, fn, line, _ := runtime.Caller(1)
	formatter := &easy.Formatter{

		TimestampFormat: "2006-01-02 15:04:05",
		LogFormat:       "%time% [%lvl%] - %msg% ",
	}
	log.SetFormatter(formatter)
}
func initPrefixFormat() {
	formatter := &prefixed.TextFormatter{
		DisableColors: true,
		ForceColors:     true,
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
		ForceFormatting: true,
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

func init() {
	initLogFormat()
	//initEasy()
	//initPrefixFormat()
}
func main() {
	//log.Info("After show format time 738")
	//fmt.Println("")
	format("Message test line and func")
	//pattern := "time=[1] level=[2] caller=[3] msg=[4]"
	//regex(pattern)
}

func format(msg string) {
	_, fn, line, _ := runtime.Caller(1)
	log.Infof("%s:%d - %s", fn, line, msg)
}

const str = "Some strings. Price: 100$. Some strings123"

func regex(pattern string) {
	re := regexp.MustCompile(`\s*(\d*)`)
	match := re.FindStringSubmatch(pattern)
	if match != nil {
		fmt.Println(match[1])
	} else {
		fmt.Println("No match!")
	}
}
