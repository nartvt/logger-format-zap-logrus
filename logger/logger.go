package logger

import (
	log "github.com/sirupsen/logrus"
	"runtime"
)

func Debug(args ...interface{}) {
	_, fn, line, _ := runtime.Caller(1)
	log.Debug("[", fn, ":", line, "] ", args)
}

func Infof(msg string) {
	_, fn, line, _ := runtime.Caller(1)
	log.Infof("%s:%d - %s", fn, line, msg)
}

func Debugf(msg string) {
	_, fn, line, _ := runtime.Caller(1)
	log.Debugf("%s:%d - %s", fn, line, msg)
}

func ErrorfOnError(msg string, err error) {
	if err != nil {
		_, fn, line, _ := runtime.Caller(1)
		log.Errorf("%s:%d %v - %s", fn, line, err, msg)
	}
}

func WarnfOnError(msg string, err error) {
	if err != nil {
		_, fn, line, _ := runtime.Caller(1)
		log.Warnf("%s:%d %v - %s", fn, line, err, msg)
	}
}

func FatalfOnError(msg string, err error) {
	if err != nil {
		_, fn, line, _ := runtime.Caller(1)
		log.Fatalf("%s:%d %v - %s", fn, line, err, msg)
	}
}

func PanicfOnError(msg string, err error) {
	if err != nil {
		_, fn, line, _ := runtime.Caller(1)
		log.Panicf("%s:%d %v - %s", fn, line, err, msg)
	}
}
