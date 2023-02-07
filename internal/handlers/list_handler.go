package handlers

import (
	"github.com/Imomali1/todo-app/config"
	"github.com/Imomali1/todo-app/pkg/cache"
)

// ListHandlerConfig ...
type ListHandlerConfig struct {
	Conf  config.Configs
	Cache cache.ICache
}

// ListHandler ...
type ListHandler struct {
	config config.Configs
	cache  cache.ICache
}

// NewListHandler ...
func NewListHandler(config *ListHandlerConfig) *ListHandler {
	return &ListHandler{
		config: config.Conf,
		cache:  config.Cache,
	}
}
