package main

import (
	"fakeApi/db"
	"fakeApi/personGen"
	log "github.com/sirupsen/logrus"
)

func main() {
	var c Config
	c.Init()

	log.SetLevel(c.ApiLogLvl)
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})

	database := db.DB{FileName: c.DBName}
	connect := database.Connect(c.DBLogLvl)

	db.DBC = connect
	personGen.DBC = connect

	db.Init()
	db.FillSampleData()

	server := NewApiServer(c.Host, c.Port)
	err := server.Run()
	if err != nil {
		log.Fatal(err)
	}
}
