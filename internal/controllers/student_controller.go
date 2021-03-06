package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/tutti2612/smart-sr-server/internal/database"
	"github.com/tutti2612/smart-sr-server/internal/models"
	"log"
	"net/http"
)

type StudentController struct{}

func (sc *StudentController) Index(c *gin.Context) {
	db := database.Connection()
	var students models.Students

	name := c.Query("name")
	classroom := c.Query("classroom")

	tx := db.Where(models.Student{Classroom: classroom})
	if name != "" {
		tx.Where("name LIKE ?", "%"+name+"%")
	}
	tx.Find(&students)

	c.JSON(http.StatusOK, students)
}

func (sc *StudentController) Show(c *gin.Context) {
	db := database.Connection()
	id := c.Param("id")
	var student models.Student
	db.First(&student, id)
	c.JSON(http.StatusOK, student)
}

func (sc *StudentController) Create(c *gin.Context) {
	//TODO バリデーションチェック
	db := database.Connection()
	var student models.Student
	if err := c.BindJSON(&student); err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}
	if result := db.Create(&student); result.Error != nil {
		log.Print(result.Error)
		c.JSON(http.StatusInternalServerError, result.Error)
		return
	}
	c.JSON(http.StatusCreated, student)
}

func (sc *StudentController) Update(c *gin.Context) {
	//TODO バリデーションチェック
	db := database.Connection()
	id := c.Param("id")
	var student models.Student
	db.First(&student, id)
	var data models.Student
	if err := c.BindJSON(&data); err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}
	if result := db.Model(&student).Updates(&data); result.Error != nil {
		log.Print(result.Error)
		c.JSON(http.StatusInternalServerError, result.Error)
		return
	}
	c.JSON(http.StatusOK, student)
}

func (sc *StudentController) Delete(c *gin.Context) {
	db := database.Connection()
	id := c.Param("id")
	db.Delete(&models.Student{}, id)
	c.JSON(http.StatusNoContent, nil)
}
