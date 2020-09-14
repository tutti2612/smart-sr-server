package models

import (
	"gorm.io/gorm"
	"time"
)

type Student struct {
	gorm.Model
	Name      string     `json:"name" gorm:"size:255"`
	Classroom string     `json:"classroom" gorm:"size:255"`
	Address   string     `json:"address" gorm:"size:255"`
	Club      string     `json:"club" gorm:"size:255"`
	Sex       string     `json:"sex" gorm:"size:1;not null;default:'0'"`
	Height    uint16     `json:"height"`
	Weight    uint16     `json:"weight"`
	Age       uint8      `json:"age"`
	Tel       string     `json:"tel" gorm:"size:255"`
	Email     string     `json:"email" gorm:"unique"`
	Birthday  *time.Time `json:"birthday" gorm:"type:date"`
}

type Students []Student
