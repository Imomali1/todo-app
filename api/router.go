package api

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/Imomali1/todo-app/config"
	"github.com/Imomali1/todo-app/internal/item"
	"github.com/Imomali1/todo-app/internal/list"
	"github.com/Imomali1/todo-app/internal/user"
	"github.com/Imomali1/todo-app/pkg/cache"
	"github.com/Imomali1/todo-app/pkg/middleware"
)

type Option struct {
	Conf  config.Configs
	Cache cache.Cache
}

type Handler struct {
	UserHandler user.Handler
	ListHandler list.Handler
	ItemHandler item.Handler
}

func NewRouter(option Option) *gin.Engine {
	gin.DefaultWriter = io.MultiWriter(os.Stdout)
	//gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	router.Use(gin.Recovery())
	router.Use(gin.Logger())
	router.Use(cors.Default())

	h := Handler{
		UserHandler: *user.NewHandler(&user.Options{
			Conf:  option.Conf,
			Cache: option.Cache,
		}),
		ListHandler: *list.NewHandler(&list.Options{
			Conf:  option.Conf,
			Cache: option.Cache,
		}),
		ItemHandler: *item.NewHandler(&item.Options{
			Conf:  option.Conf,
			Cache: option.Cache,
		}),
	}

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Welcome to Imomali's TODO Web App!")
	})
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Pong!")
	})

	api := router.Group("/api")
	{
		api.POST("/user/auth/register", h.UserHandler.Register)
		api.POST("/user/auth/login", h.UserHandler.Login)
	}

	authorized := router.Group("/api/v1")
	authorized.Use(middleware.IsAuthorized(option.Cache))
	{
		lists := authorized.Group("/lists")
		{
			lists.POST("/", h.ListHandler.CreateList)
			lists.GET("/", h.ListHandler.DetailLists)
			lists.GET("/:id", h.ListHandler.DetailListById)
			lists.PUT("/:id", h.ListHandler.UpdateList)
			lists.DELETE("/:id", h.ListHandler.DeleteList)

			items := lists.Group(":id/items")
			{
				items.POST("/", h.ItemHandler.CreateItem)
				items.GET("/", h.ItemHandler.DetailItems)
				items.GET("/:item_id", h.ItemHandler.DetailItemById)
				items.PUT("/:item_id", h.ItemHandler.UpdateItem)
				items.DELETE("/:item_id", h.ItemHandler.DeleteItem)
			}
		}

		authorized.POST("/user/logout", h.UserHandler.Logout)
	}

	return router
}
