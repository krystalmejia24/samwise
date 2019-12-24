package restapi

import (
	"net/http"
	"time"

	"github.com/krystalmejia24/samwise/pkg/restapi/encoder"
	"github.com/krystalmejia24/samwise/pkg/samwise"
)

func NewServer(svc samwise.Service) *http.Server {
	//Init router
	mux := http.NewServeMux()

	mux.Handle("/encoders/", http.HandlerFunc(encoder.New(&svc).Handle))
	mux.Handle("/encoders", http.HandlerFunc(encoder.New(&svc).Handle))

	return &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
}
