package item

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

func (h *Handler) CreateItem(c *gin.Context) {

}

func (h *Handler) DetailItems(c *gin.Context) {

}

func (h *Handler) DetailItemById(c *gin.Context) {

}

func (h *Handler) UpdateItem(c *gin.Context) {

}

func (h *Handler) DeleteItem(c *gin.Context) {

}
