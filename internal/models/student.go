package models

import (
	"time"
)

type Student struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	Name      string     `json:"name" gorm:"size:255"`
	Classroom string     `json:"classroom" gorm:"size:255"`
	Address   string     `json:"address" gorm:"size:255"`
	Club      string     `json:"club" gorm:"size:255"`
	Sex       string     `json:"sex" gorm:"size:1;not null;default:'0'"`
	Height    uint16     `json:"height,string"`
	Weight    uint16     `json:"weight,string"`
	Age       uint8      `json:"age,string"`
	Tel       string     `json:"tel" gorm:"size:255"`
	Email     string     `json:"email" gorm:"unique"`
	Birthday  *time.Time `json:"birthday" gorm:"type:date"`
}

type Students []Student
