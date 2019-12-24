package main

import (
	"encoding/base64"
	"fmt"

	"github.com/krystalmejia24/samwise/pkg/db"
	"github.com/krystalmejia24/samwise/pkg/restapi"
	"github.com/krystalmejia24/samwise/pkg/samwise"
	"github.com/krystalmejia24/samwise/pkg/scte35"

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

	//scte35 test
	b64 := "/DA0AAAAAAAA///wBQb+cr0AUAAeAhxDVUVJSAAAjn/PAAGlmbAICAAAAAAsoKGKNAIAmsnRfg=="
	byteSignal, e := base64.StdEncoding.DecodeString(b64)
	if e != nil {
		fmt.Printf("Error%v\n", e)
	}
	fmt.Println(scte35.NewScte35(byteSignal))

	log.Fatal(server.ListenAndServe())
}
