package http

import (
	"github.com/gin-gonic/gin"
)

func AuthenticationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		UserId, _ := c.Get("user_id")
		if UserId == nil {
			c.String(401, "Unauthorized")
			c.Abort()
			return
		} else {
			c.Next()
		}
	}
}
