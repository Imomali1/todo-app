package main

import (
	"context"
	"errors"
	"github.com/Imomali1/todo-app/core"
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
	conf := config.Load()

	db, err := core.InitDB(*conf)
	if err != nil {
		log.Errorf("Cannot connect to database: %s\n", err.Error())
	}

	cache := redis.NewClient(*conf)

	option := api.Option{
		Conf:  *conf,
		Cache: *cache,
		DB:    *db,
	}

	srv := server.NewServer(&option.Conf, api.NewRouter(option))

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
