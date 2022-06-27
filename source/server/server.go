package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Serve() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	router.Run(":8080")
}
