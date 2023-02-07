package api

import (
	"github.com/Imomali1/todo-app/config"
	"github.com/Imomali1/todo-app/core"
	"github.com/Imomali1/todo-app/internal/handlers"
	"github.com/Imomali1/todo-app/pkg/cache"
	"github.com/Imomali1/todo-app/pkg/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"os"
)

type Options struct {
	Conf  config.Configs
	Cache cache.ICache
	DB    core.PostgresDB
}

func InitRoutes(option *Options) *gin.Engine {
	gin.DefaultWriter = os.Stdout
	gin.DefaultErrorWriter = os.Stderr
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(cors.Default())

	h := &Handler{
		AuthHandler: *handlers.NewAuthHandler(&handlers.AuthHandlerConfig{
			Conf:  option.Conf,
			Cache: option.Cache,
		}),
		ListHandler: *handlers.NewListHandler(&handlers.ListHandlerConfig{
			Conf:  option.Conf,
			Cache: option.Cache,
		}),
		ItemHandler: *handlers.NewItemHandler(&handlers.ItemHandlerConfig{
			Conf:  option.Conf,
			Cache: option.Cache,
		}),
	}

	// No auth required endpoints
	api := router.Group("/api")
	{
		api.POST("/auth/register", h.AuthHandler.Register)
		api.POST("/auth/login", h.AuthHandler.Login)
	}

	// authorized endpoints
	authorized := router.Group("/api/v1")
	authorized.Use(middleware.IsAuthorized(option.Cache))
	{
		authorized.POST("/user/lists", h.ListHandler.CreateList)
		authorized.GET("/user/lists", h.ListHandler.GetLists)
		authorized.GET("/user/lists/:id", h.ListHandler.GetListById)
		authorized.PUT("/user/lists/:id", h.ListHandler.UpdateList)
		authorized.DELETE("/user/lists/:id", h.ListHandler.DeleteList)

		authorized.POST("/user/lists/:id/items", h.ItemHandler.CreateItem)
		authorized.GET("/user/lists/:id/items", h.ItemHandler.GetItems)
		authorized.GET("/user/lists/:id/items/:item_id", h.ItemHandler.GetItem)
		authorized.PUT("/user/lists/:id/items/:item_id", h.ItemHandler.UpdateItem)
		authorized.DELETE("/user/lists/:id/items/:item_id", h.ItemHandler.DeleteItem)
	}

	return router
}
