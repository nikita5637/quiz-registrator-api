package logger

import (
	"context"
	"io"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type contextKey int

const (
	loggerContextKey contextKey = iota
)

var (
	defaultLevel = zap.NewAtomicLevelAt(zap.InfoLevel)
	globalLogger *zap.SugaredLogger
)

func init() {
	SetGlobalLogger(NewLogger(defaultLevel, os.Stdout))
}

// NewLogger ...
func NewLogger(level zapcore.LevelEnabler, w io.Writer, options ...zap.Option) *zap.SugaredLogger {
	return NewWithSink(level, w, options...)
}

// NewWithSink ...
func NewWithSink(level zapcore.LevelEnabler, sink io.Writer, options ...zap.Option) *zap.SugaredLogger {
	if level == nil {
		level = defaultLevel
	}

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zapcore.EncoderConfig{
			TimeKey:        "ts",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			MessageKey:     "message",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		}),
		zapcore.AddSync(sink),
		level,
	)

	return zap.New(core, options...).Sugar()
}

// GetGlobalLogger ...
func GetGlobalLogger() *zap.SugaredLogger {
	return globalLogger
}

// SetGlobalLogger ...
func SetGlobalLogger(l *zap.SugaredLogger) {
	globalLogger = l
}

// FromContext ...
func FromContext(ctx context.Context) *zap.SugaredLogger {
	l := globalLogger

	if logger, ok := ctx.Value(loggerContextKey).(*zap.SugaredLogger); ok {
		l = logger
	}

	return l
}

// ToContext ...
func ToContext(ctx context.Context, l *zap.SugaredLogger) context.Context {
	return context.WithValue(ctx, loggerContextKey, l)
}

// Level ...
func Level() zapcore.Level {
	return defaultLevel.Level()
}

// SetLevel ...
func SetLevel(l zapcore.Level) {
	defaultLevel.SetLevel(l)
}

// Debug ...
func Debug(ctx context.Context, args ...interface{}) {
	FromContext(ctx).Debug(args...)
}

// Debugf ...
func Debugf(ctx context.Context, format string, args ...interface{}) {
	FromContext(ctx).Debugf(format, args...)
}

// DebugKV ...
func DebugKV(ctx context.Context, message string, kvs ...interface{}) {
	FromContext(ctx).Debugw(message, kvs...)
}

// Info ...
func Info(ctx context.Context, args ...interface{}) {
	FromContext(ctx).Info(args...)
}

// Infof ...
func Infof(ctx context.Context, format string, args ...interface{}) {
	FromContext(ctx).Infof(format, args...)
}

// InfoKV ...
func InfoKV(ctx context.Context, message string, kvs ...interface{}) {
	FromContext(ctx).Infow(message, kvs...)
}

// Warn ...
func Warn(ctx context.Context, args ...interface{}) {
	FromContext(ctx).Warn(args...)
}

// Warnf ...
func Warnf(ctx context.Context, format string, args ...interface{}) {
	FromContext(ctx).Warnf(format, args...)
}

// WarnKV ...
func WarnKV(ctx context.Context, message string, kvs ...interface{}) {
	FromContext(ctx).Warnw(message, kvs...)
}

// Error ...
func Error(ctx context.Context, args ...interface{}) {
	FromContext(ctx).Error(args...)
}

// Errorf ...
func Errorf(ctx context.Context, format string, args ...interface{}) {
	FromContext(ctx).Errorf(format, args...)
}

// ErrorKV ...
func ErrorKV(ctx context.Context, message string, kvs ...interface{}) {
	FromContext(ctx).Errorw(message, kvs...)
}

// Fatal ...
func Fatal(ctx context.Context, args ...interface{}) {
	FromContext(ctx).Fatal(args...)
}

// Fatalf ...
func Fatalf(ctx context.Context, format string, args ...interface{}) {
	FromContext(ctx).Fatalf(format, args...)
}

// FatalKV ...
func FatalKV(ctx context.Context, message string, kvs ...interface{}) {
	FromContext(ctx).Fatalw(message, kvs...)
}

// Panic ...
func Panic(ctx context.Context, args ...interface{}) {
	FromContext(ctx).Panic(args...)
}

// Panicf ...
func Panicf(ctx context.Context, format string, args ...interface{}) {
	FromContext(ctx).Panicf(format, args...)
}

// PanicKV ...
func PanicKV(ctx context.Context, message string, kvs ...interface{}) {
	FromContext(ctx).Panicw(message, kvs...)
}
