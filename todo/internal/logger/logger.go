package logger

import (
	"fmt"
	"github.com/rs/zerolog"
	"os"
)

type Logger struct {
	logger zerolog.Logger
}

func New(level string) *Logger {
	logLevel, err := zerolog.ParseLevel(level)
	if err != nil {
		logLevel = zerolog.InfoLevel
	}

	return &Logger{
		logger: zerolog.New(os.Stdout).Level(logLevel).With().Timestamp().Logger(),
	}
}

func (l *Logger) Info(args ...interface{}) {
	l.logger.Info().Msg(fmt.Sprint(args...))
}

func (l *Logger) Error(args ...interface{}) {
	l.logger.Error().Msg(fmt.Sprint(args...))
}

func (l *Logger) Fatal(args ...interface{}) {
	l.logger.Fatal().Msg(fmt.Sprint(args...))
}
