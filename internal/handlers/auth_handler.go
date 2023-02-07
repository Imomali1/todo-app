package handlers

import (
	"github.com/Imomali1/todo-app/config"
	"github.com/Imomali1/todo-app/pkg/cache"
)

// AuthHandlerConfig ...
type AuthHandlerConfig struct {
	Conf  config.Configs
	Cache cache.ICache
}

// AuthHandler ...
type AuthHandler struct {
	config config.Configs
	cache  cache.ICache
}

// NewAuthHandler ...
func NewAuthHandler(config *AuthHandlerConfig) *AuthHandler {
	return &AuthHandler{
		config: config.Conf,
		cache:  config.Cache,
	}
}
