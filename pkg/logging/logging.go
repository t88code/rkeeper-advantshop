package logging

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

var loggers = make(map[string]*logrus.Logger)

func GetLogger(loggerName string) *logrus.Logger {
	if logger, ok := loggers[loggerName]; ok {
		return logger
	}
	return nil
}

func NewLogger(debug bool, filename string, prefix string) (*logrus.Logger, error) {
	if _, err := os.Stat("./logs"); os.IsNotExist(err) {
		err := os.MkdirAll("logs", 0770)
		if err != nil {
			return nil, err
		}
	}
	file, err := os.OpenFile(
		fmt.Sprintf("logs/%s", filename),
		os.O_CREATE|os.O_APPEND|os.O_WRONLY,
		0640)
	if err != nil {
		return nil, err
	}
	logger := logrus.New()
	logger.Out = io.MultiWriter(file, os.Stdout)
	logger.Formatter = &logrus.TextFormatter{
		ForceColors:               true,
		DisableColors:             false,
		ForceQuote:                false,
		DisableQuote:              false,
		EnvironmentOverrideColors: false,
		DisableTimestamp:          false,
		FullTimestamp:             true,
		TimestampFormat:           "2006-01-02 15:04:05",
		DisableSorting:            false,
		SortingFunc:               nil,
		DisableLevelTruncation:    false,
		PadLevelText:              false,
		QuoteEmptyFields:          false,
		FieldMap:                  nil,
		CallerPrettyfier:          nil,
	}

	logger.WithField("prefix", prefix)
	if debug {
		logger.Level = logrus.DebugLevel
	} else {
		logger.Level = logrus.InfoLevel
	}

	return logger, nil
}
