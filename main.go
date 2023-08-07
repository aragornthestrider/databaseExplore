package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"go.elastic.co/ecszap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	logger, err := NewLogger("info")
	if err != nil {
		log.Fatal("Error in setting up logger")
		return
	}

	go func() {
		exitSignal := <-exit
		logger.Info("Service shutting down main due to: " + exitSignal.String())
		cancel()
	}()
}

func NewLogger(configLogLevel string) (*zap.Logger, error) {
	encoderConfig := ecszap.NewDefaultEncoderConfig()
	logLevel := parseLogLevel(configLogLevel)
	core := ecszap.NewCore(encoderConfig, os.Stdout, logLevel)
	logger := zap.New(core, zap.AddCaller())
	return logger, nil
}

func parseLogLevel(logLevel string) zapcore.Level {
	switch logLevel {
	case "debug":
		return zap.DebugLevel
	case "info":
		return zap.InfoLevel
	case "warn":
		return zap.WarnLevel
	case "error":
		return zap.ErrorLevel
	case "panic":
		return zap.PanicLevel
	case "fatal":
		return zap.FatalLevel
	default:
		return zap.InfoLevel
	}
}
