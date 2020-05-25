package main

import (
	"fmt"
	"os"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/krystalmejia24/samwise"
	"github.com/krystalmejia24/samwise/db"
	"github.com/krystalmejia24/samwise/restapi"
	"github.com/rs/zerolog"

	log "github.com/sirupsen/logrus"
)

//Config holds values for configuring a samwise server
type Config struct {
	Port     string        `envconfig:"HTTP_PORT" default:":8080"`
	Env      string        `envconfig:"ENV" default:"local"`
	LogLevel string        `envconfig:"LOG_LEVEL" default:"debug"`
	DBConn   string        `envconfig:"DB_CONNECTION" default:"redis-server:6379"`
	Timeout  time.Duration `envconfig:"TIMEOUT" default:"5s"`
}

func main() {
	//load config
	cfg, err := loadConfig()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(cfg)

	//init logger from zerolog
	logger := getLogger(cfg.LogLevel)

	//init redis db
	db := db.NewRedis(cfg.DBConn)

	//init rest api
	apiCfg := restapi.Config{
		Port:    cfg.Port,
		Timeout: cfg.Timeout,
		Svc:     *samwise.NewSvc(db, logger),
	}

	//start server
	server := restapi.NewServer(apiCfg)

	logger.Info().Str("port", cfg.Port).Msg("Starting Samwise")
	log.Fatal(server.ListenAndServe())
}

func loadConfig() (Config, error) {
	var c Config
	err := envconfig.Process("samwise", &c)
	if err != nil {
		return c, err
	}

	return c, nil
}

func getLogger(l string) zerolog.Logger {
	level, err := zerolog.ParseLevel(l)
	if err != nil || level == zerolog.NoLevel {
		level = zerolog.DebugLevel
	}

	return zerolog.New(os.Stderr).
		With().
		Timestamp().
		Logger().
		Level(level)
}
