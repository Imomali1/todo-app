package cli

import (
	"context"
	"github.com/Imomali1/todo-app/api"
	"github.com/Imomali1/todo-app/config"
	redis "github.com/Imomali1/todo-app/pkg/cache"
	"github.com/Imomali1/todo-app/pkg/server"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run() {
	log.SetFormatter(new(log.JSONFormatter))

	conf := config.Load()

	cache := redis.NewClient(conf)

	opt := &api.Options{
		Conf:  conf,
		Cache: *cache,
	}

	srv := server.NewServer(conf, api.InitRoutes(opt))

	go func() {
		if err := srv.Run(); err != nil {
			log.Errorf("Cannot run server: %s\n", err.Error())
		}
	}()

	log.Info("Server started...")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	const timeout = 5 * time.Second

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	if err := srv.Stop(ctx); err != nil {
		log.Error("failed to stop server: " + err.Error())
	}
}
