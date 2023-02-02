package api

import (
	"github.com/Imomali1/todo-app/config"
	"github.com/Imomali1/todo-app/internal/user"
	"github.com/Imomali1/todo-app/pkg"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

type Option struct {
	Conf config.Configs
}

type Handler struct {
	UserHandler user.Handler
}

func NewRouter(option Option) *gin.Engine {
	gin.DefaultWriter = io.MultiWriter(os.Stdout)
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	router.Use(gin.Recovery())
	router.Use(gin.Logger())
	router.Use(cors.Default())

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Welcome to Imomali's TODO Web App!")
	})
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Pong!")
	})

	api := router.Group("/api")
	{
		api.POST("/auth/register")
		api.POST("/auth/login")
		api.POST("/auth/logout")
	}

	authorized := router.Group("/api/v1")
	authorized.Use(pkg.IsAuthorized(option))
	{
		authorized.POST("/lists/list")
		authorized.GET("/lists")
		authorized.GET("/lists/:id")
		authorized.PUT("/lists/update/:id")
		authorized.DELETE("/lists/delete/:id")

		authorized.POST("/user/")
	}

	return router
}
