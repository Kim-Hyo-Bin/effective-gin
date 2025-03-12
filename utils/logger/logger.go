package logger

import (
	"fmt"
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

type Fields = logrus.Fields

func NewLogger(logPath string) *logrus.Logger {
	logger := logrus.New()

	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2025-03-10 15:04:05",
		ForceColors:     false,
		DisableColors:   true,
		ForceQuote:      true,
		// DisableQuote:            false,
		// EnvironmentOverrideColors: false,
		DisableLevelTruncation: true,
		// PadLevelText:            false,
		// QuoteEmptyFields:        false,
		// FieldMap:                nil,
		// SortingFunc:             nil,
		// DisableTimestamp:        false,
	})

	logFile, err := os.OpenFile(logPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Failed to open log file:", err)
		logger.SetOutput(os.Stdout)
	} else {
		multiWriter := io.MultiWriter(os.Stdout, logFile)
		logger.SetOutput(multiWriter)
	}
	logger.SetLevel(logrus.DebugLevel)

	return logger
}
