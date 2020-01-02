package main

import (
	"github.com/gsxhnd/gecko/cmd"
	"github.com/gsxhnd/gecko/logger"
	"go.uber.org/zap"
)

func main() {
	cmd.Execute()
	logger.SetLevel(logger.DebugLevel)
	logger.Debug("Debug", zap.String("123", "123"), zap.String("123", "123"))
	logger.Info("Info")
	logger.Warn("warn")
	logger.Error("Error")
	logger.DPanic("Panic")
	logger.Panic("Panic")
	logger.Fatal("Fatal")
	logger.Info("Info")
}
