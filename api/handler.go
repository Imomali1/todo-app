package api

import "github.com/Imomali1/todo-app/internal/handlers"

type Handler struct {
	AuthHandler handlers.AuthHandler
	ListHandler handlers.ListHandler
	ItemHandler handlers.ItemHandler
}
