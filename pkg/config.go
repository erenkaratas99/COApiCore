package pkg

import (
	"github.com/erenkaratas99/COApiCore/pkg/customErrors"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"time"
)

type AppConfigs struct {
	Env            string
	MongoDuration  time.Duration
	MongoClientURI string
	DBName         string
	ColName        string
	Port           string
	BaseUrl        string
}

var cfgs = map[string]AppConfigs{
	"OrderService": {
		Env:            "dev",
		MongoDuration:  time.Second * 10,
		MongoClientURI: "mongodb://localhost:27017",
		DBName:         "OrderDB",
		ColName:        "order",
		Port:           ":8000",
		BaseUrl:        "http://127.0.0.1",
	},
	"CustomerService": {
		Env:            "dev",
		MongoDuration:  time.Second * 10,
		MongoClientURI: "mongodb://localhost:27017",
		DBName:         "CustomerDB",
		ColName:        "customer",
		Port:           ":8001",
		BaseUrl:        "http://127.0.0.1",
	},
}

func GetAppConfig(app string) (*AppConfigs, error) {
	config, isExist := cfgs[app]
	if !isExist {
		return nil, customErrors.NewHTTPError(http.StatusInternalServerError,
			"ConfigErr",
			"App configs could not have fetched correctly.")
	}
	return &config, nil
}

func LogrusConfig() {
	log.SetOutput(os.Stdout)
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.InfoLevel)
}
