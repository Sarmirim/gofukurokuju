package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Connect(c *gin.RouterGroup) {
	c.GET("/ping", ping)
	c.POST("/post", post)
}

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, "pong")
}
