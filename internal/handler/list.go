package list

import (
	"github.com/Imomali1/todo-app/config"
	"github.com/Imomali1/todo-app/core"
	"github.com/Imomali1/todo-app/pkg/cache"
	"github.com/gin-gonic/gin"
	"net/http"
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

func (h *Handler) CreateList(c *gin.Context) {
	var in todolist
	if err := c.BindJSON(&in); err != nil {
		abort(c, http.StatusUnprocessableEntity, "Not valid todo list input!")
		return
	}

	if err :=
}

func (h *Handler) DetailLists(c *gin.Context) {

}

func (h *Handler) DetailListById(c *gin.Context) {

}

func (h *Handler) UpdateList(c *gin.Context) {

}

func (h *Handler) DeleteList(c *gin.Context) {

}
