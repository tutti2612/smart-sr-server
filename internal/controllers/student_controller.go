package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/tutti2612/smart-sr-server/internal/database"
	"github.com/tutti2612/smart-sr-server/internal/models"
	"net/http"
)

func Index(c *gin.Context) {
	db := database.Connection()
	var students models.Students
	db.Find(&students)
	c.JSON(http.StatusOK, students)
}

func Show(c *gin.Context) {
	db := database.Connection()
	id := c.Param("id")
	var student models.Student
	db.First(&student, id)
	c.JSON(http.StatusOK, student)
}

func Create(c *gin.Context) {
	db := database.Connection()
	var student models.Student
	c.BindJSON(&student)
	db.Create(&student)
	c.JSON(http.StatusCreated, student)
}

func Update(c *gin.Context) {
	db := database.Connection()
	id := c.Param("id")
	var student models.Student
	db.First(&student, id)
	var data models.Student
	c.BindJSON(&data)
	db.Model(&student).Updates(&data)
	c.JSON(http.StatusOK, student)
}

func Delete(c *gin.Context) {
	db := database.Connection()
	id := c.Param("id")
	db.Delete(&models.Student{}, id)
	c.JSON(http.StatusNoContent, nil)
}
