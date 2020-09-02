package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tutti2612/smart-sr-server/internal/database"
	"github.com/tutti2612/smart-sr-server/internal/model"
)

func Run() {
	db := database.Connection()
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "smart-sr-server",
		})
	})

	r.GET("/student/:id", func(c *gin.Context) {
		id := c.Param("id")
		var student model.Student
		db.First(&student, id)
		c.JSON(http.StatusOK, student)
	})

	r.GET("/students", func(c *gin.Context) {
		var students model.Students
		db.Find(&students)
		c.JSON(http.StatusOK, students)
	})

	r.POST("/student", func(c *gin.Context) {
		var student model.Student
		c.BindJSON(&student)
		db.Create(&student)
		c.JSON(http.StatusCreated, student)
	})

	r.PUT("/student/:id", func(c *gin.Context) {
	})

	r.DELETE("/student/:id", func(c *gin.Context) {
		id := c.Param("id")
		db.Delete(&model.Student{}, id)
		c.JSON(http.StatusNoContent, nil)
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
