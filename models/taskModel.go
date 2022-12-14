package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	ID     int    `gorm:"PRIMARY_KEY;AUTO_INCREMENT;NOT NULL"`
	Title  string `json:"title"`
	Active bool   `json:"active"`
}
