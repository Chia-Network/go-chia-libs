package rpcinterface

import (
	"log/slog"
	"os"
)

// SlogError returns a text handler preconfigured to ERROR log level
func SlogError() slog.Handler {
	opts := &slog.HandlerOptions{
		Level: slog.LevelError,
	}

	handler := slog.NewTextHandler(os.Stdout, opts)
	return handler
}

// SlogWarn returns a text handler preconfigured to WARN log level
func SlogWarn() slog.Handler {
	opts := &slog.HandlerOptions{
		Level: slog.LevelWarn,
	}

	handler := slog.NewTextHandler(os.Stdout, opts)
	return handler
}

// SlogInfo returns a text handler preconfigured to INFO log level
func SlogInfo() slog.Handler {
	opts := &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}

	handler := slog.NewTextHandler(os.Stdout, opts)
	return handler
}

// SlogDebug returns a text handler preconfigured to DEBUG log level
func SlogDebug() slog.Handler {
	opts := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}

	handler := slog.NewTextHandler(os.Stdout, opts)
	return handler
}
