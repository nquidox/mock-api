package db

import (
	"fakeApi/personGen"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DB struct {
	FileName string
}

var DBC *gorm.DB

func (d *DB) Connect(logLevel logger.LogLevel) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(d.FileName), &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	})

	if err != nil {
		panic(err)
	}
	return db
}

func Init() {
	err := DBC.AutoMigrate(&personGen.Person{})
	if err != nil {
		panic(err)
	}
}
