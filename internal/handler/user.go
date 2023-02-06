package user

import (
	"github.com/Imomali1/todo-app/config"
	"github.com/Imomali1/todo-app/core"
	"github.com/Imomali1/todo-app/pkg/cache"
	"github.com/gin-gonic/gin"
)

type Options struct {
	Conf  config.Configs
	Cache cache.Cache
	DB    core.PostgresClient
}

type Handler struct {
	Options
}

func NewHandler(options *Options) *Handler {
	return &Handler{
		Options: *options,
	}
}

func (h *Handler) Register(c *gin.Context) {

}

func (h *Handler) Login(c *gin.Context) {

}

func (h *Handler) Logout(c *gin.Context) {

}
