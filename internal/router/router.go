package router

import (
	"github.com/tutti2612/smart-sr-server/internal/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "smart-sr-router",
		})
	})

	r.GET("/students", controllers.Index)
	r.GET("/student/:id", controllers.Show)
	r.POST("/student", controllers.Create)
	r.PUT("/student/:id", controllers.Update)
	r.DELETE("/student/:id", controllers.Delete)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
