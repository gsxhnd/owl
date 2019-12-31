package main

import (
	"github.com/gsxhnd/gecko/cmd"
	"github.com/gsxhnd/gecko/logger"
	"go.uber.org/zap"
)

func main() {
	cmd.Execute()
	logger.SetLevel(logger.DebugLevel)
	logger.Debug("Debug")
	//logger.Info("Info")
	//logger.Warn("warn")
	//logger.Error("Error")
	logger.Debug("123", zap.String("123", "123"), zap.String("123", "123"))
}
