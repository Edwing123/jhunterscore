package main

import (
	"io"
	"log/slog"
)

func NewLogger(output io.Writer) *slog.Logger {
	logger := slog.New(
		slog.NewTextHandler(output, nil),
	)

	return logger
}
