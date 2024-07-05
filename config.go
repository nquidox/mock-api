package main

import (
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm/logger"
	"os"
)

type Config struct {
	Host      string
	Port      string
	DBName    string
	ApiLogLvl log.Level
	DBLogLvl  logger.LogLevel
}

func (c *Config) Init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	c.Host = getEnv("FAKEAPI_HOST")
	c.Port = getEnv("FAKEAPI_PORT")
	c.DBName = getEnv("DB_NAME")

	switch getEnv("API_LOG_LVL") {
	case "Trace":
		c.ApiLogLvl = log.TraceLevel
	case "Debug":
		c.ApiLogLvl = log.DebugLevel
	case "Info":
		c.ApiLogLvl = log.InfoLevel
	case "Warning":
		c.ApiLogLvl = log.WarnLevel
	case "Error":
		c.ApiLogLvl = log.ErrorLevel
	case "Fatal":
		c.ApiLogLvl = log.FatalLevel
	case "Panic":
		c.ApiLogLvl = log.PanicLevel
	default:
		c.ApiLogLvl = log.InfoLevel
	}

	switch getEnv("DB_LOG_LVL") {
	case "Info":
		c.DBLogLvl = logger.Info
	case "Warning":
		c.DBLogLvl = logger.Warn
	case "Error":
		c.DBLogLvl = logger.Error
	case "Silent":
		c.DBLogLvl = logger.Silent
	default:
		c.DBLogLvl = logger.Info
	}
}

func getEnv(key string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return ""
}
