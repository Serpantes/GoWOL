package config

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

type Conf struct {
	LogLevel     string
	ServerHost   string
	BroadcastMask string
}

var Config Conf

func InitConfig() {

	//Read .env files if exists
	log.Trace("Checking if env = PROD")
	if os.Getenv("env") != "PROD" {
		log.Info("not in PROD. Reading .env")
		var errEnv error
		log.Trace("Checking if .env.local exists")
		if _, err := os.Stat(".env.local"); err == nil {
			log.Trace("reading .env.local")
			errEnv = godotenv.Load(".env.local")
		} else {
			log.Trace("reading .env")
			errEnv = godotenv.Load(".env")
		}
		if errEnv != nil {
			log.Fatal("Can't read .env file. ", errEnv)
		}
	}

	//set up variables from env values
	Config.LogLevel = os.Getenv("LOG_LEVEL")
	Config.ServerHost = os.Getenv("SERVER_HOST")
	Config.BroadcastMask = os.Getenv("BROADCAST_MASK")
	log.Debug("Config inited")
}

func InitLogger() {
	log.Trace("Setting logger parameters")

	formatter := &log.TextFormatter{
		FullTimestamp: true,
	}
	formatter.TimestampFormat = "2006-01-02 15:04:05"

	log.SetFormatter(formatter)


	switch strings.ToLower(Config.LogLevel) {
	case "trace":
		log.SetLevel(log.TraceLevel)
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	}

	log.Debug("Logger set up done")
}
