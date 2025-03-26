package database

import (
	"errors"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	database *gorm.DB
)

func Open(driver string, source string) error {
	var e error
	switch driver {
	case "sqlite":
		database, e = gorm.Open(sqlite.Open(source))
	case "mysql":
		database, e = gorm.Open(mysql.Open(source))
	default:
		return errors.New("unsupported database")
	}
	if e != nil {
		return e
	}
	e = database.AutoMigrate(&UserModel{})
	return e
}
