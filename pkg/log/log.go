package log

import (
	"log"
	"os"
	"time"

	"github.com/rs/zerolog"
)

var Logger *zerolog.Logger // nolint

func init() { // nolint
	zerolog.TimeFieldFormat = time.RFC3339

	setGlobalLevel(LogLevelDefault)
	setGlobalFormat(LogFormatDefault)
	log.SetFlags(0)
	log.SetOutput(Logger)

	Logger.Info().Msgf("logger initialized with %s level and %s format output",
		LogLevelDefault.String(),
		LogFormatDefault,
	)
}

func For(source string) *zerolog.Logger {
	logger := Logger.With().Str("source", source).Logger()

	return &logger
}

func SetGlobalLevel(level LogLevel) {
	setGlobalLevel(level)

	Logger.Info().Msgf("logger level change to %s",
		level.String(),
	)
}

func setGlobalLevel(level LogLevel) {
	zerolog.SetGlobalLevel(level)
}

func SetGlobalFormat(format LogFormat) {
	setGlobalFormat(format)

	Logger.Info().Msgf("logger format output change to %s",
		string(format),
	)
}

func setGlobalFormat(format LogFormat) {
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger() // LogFormatJson

	if format == LogFormatConsole {
		logger = zerolog.New(zerolog.ConsoleWriter{
			Out:        os.Stderr,
			TimeFormat: time.Kitchen,
		}).With().Timestamp().Logger()
	}

	Logger = &logger
}
