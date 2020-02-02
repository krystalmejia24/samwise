package main

import (
	"github.com/krystalmejia24/samwise/pkg/db"
	"github.com/krystalmejia24/samwise/pkg/restapi"
	"github.com/krystalmejia24/samwise/pkg/samwise"

	log "github.com/sirupsen/logrus"
)

func main() {
	//init logger from logrus
	logger := log.New()
	logger.Level = log.InfoLevel

	//init redis db
	db := db.NewRepo()

	server := restapi.NewServer(samwise.Service{
		DB:     db,
		Logger: logger,
	})

	log.Fatal(server.ListenAndServe())
}
