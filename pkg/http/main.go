package http

import (
	auth "backend/pkg/http/auth"
	"fmt"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(auth.TokenAuthMiddleware())
	authenticated := r.Group("/", auth.AuthenticationMiddleware())
	authenticated.GET("/ping", func(c *gin.Context) {
		userId, _ := c.Get("user_id")
		c.String(200, fmt.Sprintf("%v pong\n", userId))
	})
	return r
}
