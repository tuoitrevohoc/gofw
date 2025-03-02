package main

import (
	"github.com/tuoitrevohoc/gofw/backend/packages/gofw"
	"go.uber.org/zap"
)

func main() {
	serviceManager := gofw.NewServiceManager()
	serviceManager.AddHealthServer("8000")
	logger := serviceManager.Logger()

	err := serviceManager.Run()
	if err != nil {
		logger.Error("Failed to run service manager", zap.Error(err))
	}
}
