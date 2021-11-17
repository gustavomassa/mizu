package logger

import (
	log "github.com/sirupsen/logrus"
	"io"
	"os"
)

var Log = log.New()

func InitLogger(logPath string) {
	f, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		Log.Infof("Failed to open mizu log file: %v, err %v", logPath, err)
	}

	fileLogger := log.New()
	fileLogger.SetOutput(f)
	fileWriter := fileLogger.WriterLevel(log.DebugLevel)

	consoleLogger := log.New()
	consoleLogger.SetOutput(os.Stderr)
	consoleWriter := consoleLogger.WriterLevel(log.InfoLevel)

	mw := io.MultiWriter(fileWriter, consoleWriter)
	Log.SetOutput(mw)
}

func InitLoggerStderrOnly(level log.Level) {
	Log.SetOutput(os.Stderr)
	Log.SetLevel(level)
}
