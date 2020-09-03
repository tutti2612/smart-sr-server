package models

import (
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name      string `json:"name"`
	Classroom string `json:"classroom"`
}

type Students []Student
