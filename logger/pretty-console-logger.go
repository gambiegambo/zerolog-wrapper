package logger

// From: https://github.com/rs/zerolog
// Fatal starts a new message with fatal level. The os.Exit(1) function is called by the Msg method, which terminates the program immediately.
// The panic() function is called by the Msg method, which stops the ordinary flow of a goroutine.

import (
	"fmt"
	"github.com/rs/zerolog"
	"os"
	"strings"
	"sync"
	"time"
)

var (
	prettyConsoleLoggerOnce sync.Once
	prettyConsoleLogger     zerolog.Logger
)

// New gets loggers instance
// callerLength: default is 30
// possible log level: debug, info, warn, error, fatal, panic, and disabled (ignore case)
// 2019-06-17T12:29:33.201563+07:00 | WARN  |  main.go:13                     > Test name:gam
func New(logLevel string) {
	prettyConsoleLoggerOnce.Do(func() {
		logLevelEnv := getLogLevel(logLevel)

		if logLevelEnv == zerolog.Disabled {
			prettyConsoleLogger = zerolog.Nop()
			return
		}

		zerolog.TimeFieldFormat = time.RFC3339Nano
		zerolog.SetGlobalLevel(logLevelEnv)
		output := zerolog.ConsoleWriter{
			Out:        os.Stdout,
			TimeFormat: time.RFC3339Nano,
		}

		output.FormatTimestamp = func(i interface{}) string {
			return fmt.Sprintf("%-32s", i)
		}

		output.FormatLevel = func(i interface{}) string {
			return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
		}

		output.FormatCaller = func(i interface{}) string {
			var caller string

			if cc, ok := i.(string); ok {
				caller = cc
			}

			if len(caller) > 0 {
				caller := strings.Split(i.(string), "/")
				return fmt.Sprintf(" %-30s >", caller[len(caller)-1])
			}

			return fmt.Sprintf(" %-30s >", "unknown")
		}

		output.FormatFieldName = func(i interface{}) string {
			return fmt.Sprintf("%s: ", i)
		}

		output.FormatErrFieldName = func(i interface{}) string {
			return fmt.Sprintf("%s: ", i)
		}

		prettyConsoleLogger = zerolog.New(output).With().Caller().Timestamp().Logger()
	})
}

// Get returns zerolog.Logger instance
func Get() *zerolog.Logger {
	return &prettyConsoleLogger
}

/**
getLogLevel converts input log level to zerolog level
in case error on parsing, do set log level to INFO
possible log level: debug, info, warn, error, fatal, panic, and disabled
*/
func getLogLevel(level string) zerolog.Level {
	level = strings.TrimSpace(strings.ToLower(level))

	if level == "disabled" {
		return zerolog.Disabled
	}

	parsedLevel, err := zerolog.ParseLevel(level)
	if err != nil || parsedLevel == zerolog.NoLevel {
		parsedLevel = zerolog.InfoLevel
	}

	return parsedLevel
}
