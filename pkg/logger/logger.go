package logger

import (
	"golang.org/x/exp/slog"
	"os"
)

type LoggerInterface interface {
	Info(message string, fields map[string]interface{})
	Error(message string, fields map[string]interface{})
	Debug(message string, fields map[string]interface{})
}

type CustomLogger struct {
	innerLogger *slog.Logger
}

func NewLogger() *CustomLogger {
	// Create a JSON handler that writes to stdout
	jsonHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})

	// Create a new logger instance with the JSON handler
	logger := &CustomLogger{
		innerLogger: slog.New(jsonHandler),
	}

	return logger
}

func (l *CustomLogger) Info(message string, fields map[string]interface{}) {
	attrs := toSlogAttrs(fields)
	var attrsAny []any

	attrsAny = make([]any, len(attrs))
	for i, attr := range attrs {
		attrsAny[i] = attr
	}

	l.innerLogger.Info(message, attrsAny...)
}

func (l *CustomLogger) Error(message string, fields map[string]interface{}) {
	attrs := toSlogAttrs(fields)
	var attrsAny []any

	attrsAny = make([]any, len(attrs))
	for i, attr := range attrs {
		attrsAny[i] = attr
	}

	l.innerLogger.Error(message, attrsAny...)
}

func (l *CustomLogger) Debug(message string, fields map[string]interface{}) {
	attrs := toSlogAttrs(fields)
	var attrsAny []any

	attrsAny = make([]any, len(attrs))
	for i, attr := range attrs {
		attrsAny[i] = attr
	}

	l.innerLogger.Debug(message, attrsAny...)
}

// toSlogAttrs converts a map to slog's Attr type
func toSlogAttrs(fields map[string]interface{}) []slog.Attr {
	attrs := make([]slog.Attr, 0, len(fields))
	for k, v := range fields {
		attrs = append(attrs, slog.Any(k, v))
	}
	return attrs
}
