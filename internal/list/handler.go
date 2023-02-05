package list

import (
	"github.com/Imomali1/todo-app/config"
	"github.com/Imomali1/todo-app/pkg/cache"
	"github.com/gin-gonic/gin"
)

type Options struct {
	Conf  config.Configs
	Cache cache.Cache
}

type Handler struct {
	Options
}

func NewHandler(options *Options) *Handler {
	return &Handler{
		Options: *options,
	}
}

func (h *Handler) CreateList(c *gin.Context) {

}

func (h *Handler) DetailLists(c *gin.Context) {

}

func (h *Handler) DetailListById(c *gin.Context) {

}

func (h *Handler) UpdateList(c *gin.Context) {

}

func (h *Handler) DeleteList(c *gin.Context) {

}
