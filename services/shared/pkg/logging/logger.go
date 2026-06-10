package logging

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

func New(service string) zerolog.Logger {
	output := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339,
	}

	return zerolog.New(output).
		Level(zerolog.DebugLevel).
		With().
		Timestamp().
		Str("service", service).
		Caller().
		Logger()
}

func NewJSON(service string) zerolog.Logger {
	return zerolog.New(os.Stdout).
		Level(zerolog.InfoLevel).
		With().
		Timestamp().
		Str("service", service).
		Caller().
		Logger()
}
