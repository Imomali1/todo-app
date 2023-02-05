package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/Imomali1/todo-app/api"
	"github.com/Imomali1/todo-app/config"
	redis "github.com/Imomali1/todo-app/pkg/cache"
	"github.com/Imomali1/todo-app/pkg/server"
)

func main() {
	log.Info("Before config")
	conf := config.Load()
	log.Info("After config")

	cache := redis.NewClient(conf)

	option := api.Option{
		Conf:  conf,
		Cache: *cache,
	}

	srv := server.NewServer(&option.Conf, api.NewRouter(option))

	log.Info("Server is running on ", conf.HTTP.Host+conf.HTTP.Port)

	go func() {
		if err := srv.Run(); !errors.Is(err, http.ErrServerClosed) {
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
