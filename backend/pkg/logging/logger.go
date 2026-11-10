package logging

import (
	"fmt"
	"github.com/rs/zerolog"
	"os"
	"strings"
	"time"
)

type ServerLogger struct {
	zerolog.Logger
}

var Logger ServerLogger

func NewServerLogger() ServerLogger {
	// create output configuration
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}

	// Format level: fatal, error, debug, info, warn (pretty colours)
	output.FormatLevel = func(i interface{}) string {
		level := strings.ToUpper(fmt.Sprintf("%s", i))

		switch level {
		case "TRACE":
			return "\033[32m| SUCCESS  |\033[0m"
		case "DEBUG":
			return "\033[36m| DEBUG  |\033[0m"
		case "INFO":
			return "\033[34m| INFO   |\033[0m"
		case "WARN":
			return "\033[33m| WARN   |\033[0m"
		case "ERROR":
			return "\033[31m| ERROR  |\033[0m"
		case "FATAL":
			return "\033[35m| FATAL  |\033[0m"
		default:
			return fmt.Sprintf("| %-6s|", level)
		}
	}
	output.FormatFieldName = func(i interface{}) string {
		return fmt.Sprintf("%s:", i)
	}
	output.FormatFieldValue = func(i interface{}) string {
		return fmt.Sprintf("%s", i)
	}

	// format error
	output.FormatErrFieldName = func(i interface{}) string {
		return fmt.Sprintf("%s: ", i)
	}

	zLog := zerolog.New(output).With().Timestamp().Logger()
	Logger = ServerLogger{zLog}
	return Logger
}

func (l *ServerLogger) LogInfo() *zerolog.Event {
	logger := l.Logger.With().Caller().Logger()
	return logger.Info()
}

func (l *ServerLogger) LogError() *zerolog.Event {
	logger := l.Logger.With().Caller().Logger()
	return logger.Error()
}

func (l *ServerLogger) LogDebug() *zerolog.Event {
	logger := l.Logger.With().Caller().Logger()
	return logger.Debug()
}

func (l *ServerLogger) LogWarn() *zerolog.Event {
	logger := l.Logger.With().Caller().Logger()
	return logger.Warn()
}

func (l *ServerLogger) LogFatal() *zerolog.Event {
	logger := l.Logger.With().Caller().Logger()
	return logger.Fatal()
}
