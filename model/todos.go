package model

import "gorm.io/gorm"

type Todos struct {
	gorm.Model
	ID   uint   `gorm:"primaryKey"`
	Text string `gorm:"not null"`
	Done bool   `gorm:"default:false"`
}
