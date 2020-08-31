package server

import (
	"net/http"
	
	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "smart-sr-server",
		})
	})

	r.GET("/students", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"students": "asfdasd",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
