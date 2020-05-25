package restapi

import (
	"net/http"
	"time"

	"github.com/krystalmejia24/samwise"
	"github.com/krystalmejia24/samwise/restapi/encoder"
)

//Config holds all configuration values for the server instance
type Config struct {
	Port    string
	Timeout time.Duration
	Svc     samwise.Service
}

//NewServer createa and returns a new server instance to hanndle http requests
func NewServer(c Config) *http.Server {
	//Init router
	mux := http.NewServeMux()

	mux.Handle("/encoders/", http.HandlerFunc(encoder.New(&c.Svc).Handle))
	mux.Handle("/encoders", http.HandlerFunc(encoder.New(&c.Svc).Handle))

	return &http.Server{
		Addr:         c.Port,
		Handler:      mux,
		ReadTimeout:  c.Timeout,
		WriteTimeout: c.Timeout,
		IdleTimeout:  c.Timeout,
	}
}
