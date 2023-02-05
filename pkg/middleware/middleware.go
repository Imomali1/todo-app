package middleware

import (
	"encoding/json"
	"github.com/Imomali1/todo-app/internal/user"
	"github.com/Imomali1/todo-app/pkg/cache"
	"github.com/gin-gonic/gin"
	"net/http"
)

const authHeader = "Authorization"

func IsAuthorized(cache cache.Cache) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get(authHeader)

		if token == "" {
			abortWithError(c, http.StatusUnauthorized, "Unauthorized!")
			return
		}

		keys, err := cache.Keys(c, "auth:*:"+token)
		if err != nil {
			abortWithError(c, http.StatusInternalServerError, "Cannot get keys from cache!")
			return
		}

		if len(keys) == 0 {
			abortWithError(c, http.StatusInternalServerError, "There is no key in cache!")
			return
		}

		data, err := cache.Get(c, keys[0])
		if err != nil {
			abortWithError(c, http.StatusInternalServerError, "Cannot get data from cache!")
			return
		}

		var authModel user.AuthModel
		if err := json.Unmarshal([]byte(data), &authModel); err != nil {
			abortWithError(c, http.StatusUnprocessableEntity, "Cannot parse data!")
			return
		}

		c.Set("AuthModel", authModel)
		c.Next()
	}
}

func abortWithError(ctx *gin.Context, status int, errMsg string) {
	ctx.JSON(status, gin.H{"error": errMsg})
	ctx.AbortWithStatus(status)
}
