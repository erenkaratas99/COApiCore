package pkg

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func InitLogrusConfig() {
	log.SetOutput(os.Stdout)
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.InfoLevel)
}
