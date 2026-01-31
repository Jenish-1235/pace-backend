package utils

import (
	"context"
	"fmt"

	"go.uber.org/zap"
)

type ctxKey struct{}

type StandardLogger struct{
	*zap.Logger
}

func (l *StandardLogger) Infof(format string, args ...any) {
	l.Info(fmt.Sprintf(format, args...))
}

func (l *StandardLogger) Debugf(format string, args ...any) {
	l.Debug(fmt.Sprintf(format, args...))
}

func (l *StandardLogger) Errorf(format string, args ...any) {
	l.Error(fmt.Sprintf(format, args...))
}

func (l *StandardLogger) Warnf(format string, args ...any) {
	l.Warn(fmt.Sprintf(format, args...))
}

func (l *StandardLogger) Fatalf(format string, args ...any) {
	l.Fatal(fmt.Sprintf(format, args...))
}

func (l *StandardLogger) Panicf(format string, args ...any) {
	l.Panic(fmt.Sprintf(format, args...))
}

var appLogger *StandardLogger

func InitLogger(level string, environment string) {
	var cfg zap.Config

	if environment == "production" {
		cfg = zap.NewProductionConfig()
	} else {
		cfg = zap.NewDevelopmentConfig()
	}

	logLevel := zap.InfoLevel
	if err := logLevel.Set(level); err != nil {
		logLevel = zap.InfoLevel
	}

	cfg.Level = zap.NewAtomicLevelAt(logLevel)

	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	
	appLogger = &StandardLogger{logger}
}

func GetLogger() *StandardLogger {
	if appLogger == nil {
		panic("logger not initialised")
	}
	return appLogger
}

func LoggerWithContext(ctx context.Context, logger *StandardLogger) context.Context {
	return context.WithValue(ctx, ctxKey{}, logger)
}

func LoggerFromContext(ctx context.Context) *StandardLogger {
	if logger, ok := ctx.Value(ctxKey{}).(*StandardLogger); ok {
		return logger
	}
	return GetLogger()
}