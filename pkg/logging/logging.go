package logging

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

type Logger struct {
	*logrus.Logger
}

var loggers = make(map[string]*Logger)

func GetLogger(loggerName string) (*Logger, error) {
	if l, ok := loggers[loggerName]; ok {
		return l, nil
	}
	return nil, fmt.Errorf("not found logger: %s", loggerName)
}

func NewLogger(debug bool, filename string, prefix string, loggerName string) (*Logger, error) {
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

	if logger, ok := loggers[loggerName]; ok {
		return logger, nil
	}

	l := logrus.New()
	l.Out = io.MultiWriter(file, os.Stdout)
	l.Formatter = &logrus.TextFormatter{
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
	l.WithField("prefix", prefix)
	if debug {
		l.Level = logrus.DebugLevel
	} else {
		l.Level = logrus.InfoLevel
	}
	loggers[loggerName] = &Logger{l}
	return loggers[loggerName], nil
}
