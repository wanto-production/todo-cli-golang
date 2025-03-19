package model

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	Database *gorm.DB
	err      error
)

func Init() error {
	Database, err = gorm.Open(sqlite.Open("todos.db"), &gorm.Config{})
	return err
}
