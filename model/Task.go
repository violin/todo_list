package model

import "github.com/jinzhu/gorm"

type Task struct {
	gorm.Model
	Userid      string
	Taskcontent string `gorm:"type:text"`
	Isdeleted   bool   `gorm:"default:false"`
}
