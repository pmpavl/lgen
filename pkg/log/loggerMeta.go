package log

import (
	"fmt"

	"github.com/rs/zerolog"
)

type LogLevel = zerolog.Level

const (
	LogLevelTrace LogLevel = iota - 1
	LogLevelDebug
	LogLevelInfo
	LogLevelWarn
	LogLevelError
	LogLevelFatal
	LogLevelPanic
	LogLevelNoLevel
	LogLevelDisabled

	LogLevelDefault = LogLevelDebug
)

func ParseLogLevel(level string) (LogLevel, error) {
	logLevel, err := zerolog.ParseLevel(level)
	if err != nil {
		return LogLevelDefault, fmt.Errorf("no such level: %s", level)
	}

	return logLevel, nil
}

type LogFormat = string

const (
	LogFormatJson    LogFormat = "json"
	LogFormatConsole LogFormat = "console"

	LogFormatDefault = LogFormatConsole
)

func ParseLogFormat(format string) (LogFormat, error) {
	switch format {
	case string(LogFormatJson):
		return LogFormatJson, nil
	case string(LogFormatConsole):
		return LogFormatConsole, nil
	}

	return LogFormatDefault, fmt.Errorf("no such format: %s", format)
}
