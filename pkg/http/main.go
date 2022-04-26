package http

import "github.com/gin-gonic/gin"


func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong\n")
	})
	return r
}