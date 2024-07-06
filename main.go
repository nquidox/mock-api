package main

import (
	"fakeApi/db"
	"fakeApi/personGen"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"path"
	"runtime"
)

func main() {
	var c Config
	c.Init()
	logSetup(c)

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

func logSetup(c Config) {
	log.SetFormatter(
		&log.TextFormatter{
			FullTimestamp: true,
			CallerPrettyfier: func(f *runtime.Frame) (string, string) {
				filename := path.Base(f.File)
				return fmt.Sprintf("%s()", f.Function), fmt.Sprintf(" %s:%d", filename, f.Line)
			},
		},
	)
	if l, err := log.ParseLevel("debug"); err == nil {
		log.SetLevel(l)
		log.SetReportCaller(l == log.DebugLevel)
		log.SetOutput(os.Stdout)
	}
	log.SetOutput(os.Stdout)

	log.SetLevel(c.ApiLogLvl)
}
