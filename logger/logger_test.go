package logger

import (
	"bytes"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

func Test_FormatLog(t *testing.T) {
	os.Setenv("APP_ENV", "test")
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()
	initEasyFormat()
	Infof("test log")

	timev := time.Now()
	currentFolder, _ := os.Getwd()
	assert.Equal(t, timev.Format("[2006-01-02 15:04:05]")+fmt.Sprintf("  INFO %s/logger_test.go:21 - test log\n", currentFolder), buf.String())
}
