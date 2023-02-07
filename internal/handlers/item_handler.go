package handlers

import (
	"github.com/Imomali1/todo-app/config"
	"github.com/Imomali1/todo-app/pkg/cache"
)

// ItemHandlerConfig ...
type ItemHandlerConfig struct {
	Conf  config.Configs
	Cache cache.ICache
}

// ItemHandler ...
type ItemHandler struct {
	config config.Configs
	cache  cache.ICache
}

// NewItemHandler ...
func NewItemHandler(config *ItemHandlerConfig) *ItemHandler {
	return &ItemHandler{
		config: config.Conf,
		cache:  config.Cache,
	}
}
