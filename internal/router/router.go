package router

import (
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/tutti2612/smart-sr-server/internal/controllers"
)

func Run() {
	r := gin.Default()

	// CORS 対応
	r.Use(cors.Default())

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "smart-sr-server-go",
		})
	})

	var sc controllers.StudentController
	r.GET("/students", sc.Index)
	r.GET("/student/:id", sc.Show)
	r.POST("/student", sc.Create)
	r.PUT("/student/:id", sc.Update)
	r.DELETE("/student/:id", sc.Delete)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}
