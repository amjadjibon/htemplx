package logger

import (
	"context"
	"log/slog"
	"os"
	"strings"
)

const (
	AppName = "htemplx"
)

type contextKey int

const (
	RequestIDKey contextKey = iota
	TraceIDKey
	TraceDataKey
)

type HandlerRequestID struct {
	slog.Handler
}

func (h HandlerRequestID) Handle(ctx context.Context, r slog.Record) error {
	if requestID, ok := ctx.Value(RequestIDKey).(string); ok {
		r.Add("request_id", slog.StringValue(requestID))
	}

	if traceID, ok := ctx.Value(TraceIDKey).(string); ok {
		r.Add("trace_id", slog.StringValue(traceID))
	}

	if traceData := ctx.Value(TraceDataKey); traceData != nil {
		r.Add("trace_data", traceData)
	}

	return h.Handler.Handle(ctx, r)
}

// GetLogLevel returns slog.Level by level string
func GetLogLevel(level string) slog.Level {
	var logLevel slog.Level
	switch strings.ToLower(level) {
	case "debug":
		logLevel = slog.LevelDebug
	case "info":
		logLevel = slog.LevelInfo
	case "warn":
		logLevel = slog.LevelWarn
	case "error":
		logLevel = slog.LevelError
	default:
		logLevel = slog.LevelInfo
	}
	return logLevel
}

func trimFuncName(funcName string) string {
	if funcName == "" {
		return ""
	}

	parts := strings.Split(funcName, ".")
	return parts[len(parts)-1]
}

func trimFileName(fileName string) string {
	if fileName == "" {
		return ""
	}

	index := strings.LastIndex(fileName, AppName)
	if index == -1 {
		return fileName
	}

	return fileName[index+len(AppName)+1:]
}

func Replacer(groups []string, a slog.Attr) slog.Attr {
	if a.Key == slog.SourceKey {
		source := a.Value.Any().(*slog.Source)
		source.File = trimFileName(source.File)
		source.Function = trimFuncName(source.Function)
	}
	return a
}

// GetLogger returns a slog.Logger with HandlerRequestID
// and slog.Source replaced by relative path by logger level string
func GetLogger(level string) *slog.Logger {
	options := &slog.HandlerOptions{
		Level:       GetLogLevel(level),
		AddSource:   true,
		ReplaceAttr: Replacer,
	}

	handler := HandlerRequestID{Handler: slog.NewJSONHandler(os.Stderr, options)}
	return slog.New(handler).With()
}

// SetDefault sets slog.DefaultLogger
func SetDefault(logger *slog.Logger) {
	slog.SetDefault(logger)
}
