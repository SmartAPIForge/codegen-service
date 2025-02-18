package main

import (
	"codegen-service/internal/app"
	"codegen-service/internal/config"
	"codegen-service/internal/lib/logger"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := config.MustLoad()

	log := logger.MustSetupLogger(cfg.Env)

	application := app.NewApp(
		log,
		cfg.GRPC.Port,
		cfg.RedisAddress,
		cfg.RedisDb,
	)
	application.GrpcApp.MustRun()

	stopWait(application)
}

func stopWait(application *app.App) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	<-stop
	application.GrpcApp.Stop()
	err := application.RedisClient.Close()
	if err != nil {
		fmt.Println("failed to break redis connection")
	}
}
