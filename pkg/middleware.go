package pkg

import (
	"github.com/Imomali1/todo-app/api"
	"github.com/gin-gonic/gin"
)

func IsAuthorized(option api.Option) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
