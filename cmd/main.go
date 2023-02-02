package main

import (
	"context"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Imomali1/todo-app/api"
	"github.com/Imomali1/todo-app/config"
	"github.com/Imomali1/todo-app/pkg/server"
)

func main() {
	conf := config.Load()

	option := api.Option{
		Conf: conf,
	}

	srv := server.NewServer(&option.Conf, api.NewRouter(option))

	go func() {
		if err := srv.Run(); err != nil {
			log.Error("error occurred while running http server: " + err.Error())
		}
	}()

	log.Info("Server started...")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	const timeout = 5 * time.Second

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	if err := srv.Stop(ctx); err != nil {
		log.Error("failed to stop server: " + err.Error())
	}
}
