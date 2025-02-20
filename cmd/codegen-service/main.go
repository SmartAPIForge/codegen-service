package main

import (
	"codegen-service/internal/app"
	"codegen-service/internal/config"
	_ "codegen-service/internal/kafka"
	"codegen-service/internal/lib/logger"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := config.MustLoad()

	log := logger.MustSetupLogger(cfg.Env)

	application := app.NewApp(
		log,
		cfg,
	)
	application.GrpcApp.MustRun()

	stopWait(application)
}

func stopWait(application *app.App) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	<-stop
	application.GrpcApp.Stop()
}
