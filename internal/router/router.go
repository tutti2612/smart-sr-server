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
			"message": "smart-sr-server",
		})
	})

	var sc controllers.StudentController
	r.GET("/students", sc.Index)
	r.GET("/student/:id", sc.Show)
	r.POST("/student", sc.Create)
	r.PUT("/student/:id", sc.Update)
	r.DELETE("/student/:id", sc.Delete)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
