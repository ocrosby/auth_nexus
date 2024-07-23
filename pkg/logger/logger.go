package logger

import (
	"golang.org/x/exp/slog"
	"os"
)

// Logger is the global logger instance
var Logger *slog.Logger

// Init initializes the global logger with a JSON handler
func Init() {
	// Create a JSON handler that writes to stdout
	jsonHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})

	// Create a new logger instance with the JSON handler
	Logger = slog.New(jsonHandler)
}

// Info logs an info message with structured data
func Info(message string, fields map[string]interface{}) {
	attrs := toSlogAttrs(fields)
	var attrsAny []any

	attrsAny = make([]any, len(attrs))
	for i, attr := range attrs {
		attrsAny[i] = attr
	}

	Logger.Info(message, attrsAny...)
}

// Error logs an error message with structured data
func Error(message string, fields map[string]interface{}) {
	attrs := toSlogAttrs(fields)
	var attrsAny []any

	attrsAny = make([]any, len(attrs))
	for i, attr := range attrs {
		attrsAny[i] = attr
	}

	Logger.Error(message, attrsAny...)
}

// Debug logs a debug message with structured data
func Debug(message string, fields map[string]interface{}) {
	attrs := toSlogAttrs(fields)
	var attrsAny []any

	attrsAny = make([]any, len(attrs))
	for i, attr := range attrs {
		attrsAny[i] = attr
	}

	Logger.Debug(message, attrsAny...)
}

// toSlogAttrs converts a map to slog's Attr type
func toSlogAttrs(fields map[string]interface{}) []slog.Attr {
	attrs := make([]slog.Attr, 0, len(fields))
	for k, v := range fields {
		attrs = append(attrs, slog.Any(k, v))
	}
	return attrs
}
