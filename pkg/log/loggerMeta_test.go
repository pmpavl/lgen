package log_test

import (
	"testing"

	"github.com/pmpavl/lgen/pkg/log"
	"github.com/stretchr/testify/assert"
)

type answParseLogLevel struct {
	level  string
	errStr string
}

func ttTestParseLogLevel() []struct { // nolint
	name  string
	level string
	answ  answParseLogLevel
} {
	return []struct {
		name  string
		level string
		answ  answParseLogLevel
	}{
		{
			name:  "trace level",
			level: "trace",
			answ: answParseLogLevel{
				level:  "trace",
				errStr: "",
			},
		},
		{
			name:  "debug level",
			level: "debug",
			answ: answParseLogLevel{
				level:  "debug",
				errStr: "",
			},
		},
		{
			name:  "info level",
			level: "info",
			answ: answParseLogLevel{
				level:  "info",
				errStr: "",
			},
		},
		{
			name:  "warn level",
			level: "warn",
			answ: answParseLogLevel{
				level:  "warn",
				errStr: "",
			},
		},
		{
			name:  "error level",
			level: "error",
			answ: answParseLogLevel{
				level:  "error",
				errStr: "",
			},
		},
		{
			name:  "fatal level",
			level: "fatal",
			answ: answParseLogLevel{
				level:  "fatal",
				errStr: "",
			},
		},
		{
			name:  "panic level",
			level: "panic",
			answ: answParseLogLevel{
				level:  "panic",
				errStr: "",
			},
		},
		{
			name:  "nolevel level", // empty level
			level: "",
			answ: answParseLogLevel{
				level:  "",
				errStr: "",
			},
		},
		{
			name:  "disabled level",
			level: "disabled",
			answ: answParseLogLevel{
				level:  "disabled",
				errStr: "",
			},
		},
		{
			name:  "wrong level",
			level: "wrong",
			answ: answParseLogLevel{
				level:  "debug",
				errStr: "no such level: wrong",
			},
		},
	}
}

func TestParseLogLevel(t *testing.T) {
	t.Parallel()

	tt := ttTestParseLogLevel()
	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			level, err := log.ParseLogLevel(tc.level)

			var errString string
			if err != nil {
				errString = err.Error()
			}

			answ := answParseLogLevel{level.String(), errString}

			assert.EqualValues(t, answ, tc.answ)
		})
	}
}

type answParseLogFormat struct {
	format string
	errStr string
}

func ttTestParseLogFormat() []struct {
	name   string
	format string
	answ   answParseLogFormat
} {
	return []struct {
		name   string
		format string
		answ   answParseLogFormat
	}{
		{
			name:   "console format",
			format: "console",
			answ: answParseLogFormat{
				format: "console",
				errStr: "",
			},
		},
		{
			name:   "json format",
			format: "json",
			answ: answParseLogFormat{
				format: "json",
				errStr: "",
			},
		},
		{
			name:   "empty format",
			format: "",
			answ: answParseLogFormat{
				format: "console",
				errStr: "no such format: ",
			},
		},
		{
			name:   "wrong format",
			format: "wrong",
			answ: answParseLogFormat{
				format: "console",
				errStr: "no such format: wrong",
			},
		},
	}
}

func TestParseLogFormat(t *testing.T) {
	t.Parallel()

	tt := ttTestParseLogFormat()
	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			format, err := log.ParseLogFormat(tc.format)

			var errString string
			if err != nil {
				errString = err.Error()
			}

			answ := answParseLogFormat{format, errString}

			assert.EqualValues(t, answ, tc.answ)
		})
	}
}
