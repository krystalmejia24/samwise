package encoder

import (
	"fmt"
	"net/http"
	"time"
)

type Service interface {
	InsertEvent()
	InsertScte()
	InsertID3()
}

// HTTPClient hold interface declaration of our http clients
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client holds configuration for http clients
type Client struct {
	Timeout time.Duration `envconfig:"CLIENT_TIMEOUT" default:"5s"`
	HTTPClient
}

//Encoder struct represents encoding instance persisted in the repo
type Encoder struct {
	IP     string    `json:"ip"`
	Config *Config   `json:"config,omitempty"`
	Stream *[]Stream `json:"stream,omitempty"`
}

//Config struct holds authentication needed for encoding
type Config struct {
	User   string `json:"User,omitempty"`
	APIKey string `json:"api_key,omitempty"`
}

//Stream is the event associated to an encoding instance
type Stream struct {
	ID int `json:"id,omitempty"`
}

//StatusError object sets the error response
type StatusError struct {
	Code int
	Msg  string
	body string
}

//NotFound will set the status code to 404
func (e StatusError) NotFound() bool {
	return e.Code == 404
}

//Error will return a formatted status error
func (e StatusError) Error() string {
	return fmt.Sprintf("http status: %d: %q", e.Code, e.body)
}
