package db

//Encoder struct represents encoding instance persisted in the repo
type Encoder struct {
	IP     string    `json:"ip"`
	Config *Config   `json:"config,omitempty"`
	Stream *[]Stream `json:"stream,omitempty"`
}

//Config struct holds authentication needed for encoding
type Config struct {
	Username string `json:"username,omitempty"`
	APIKey   string `json:"api_key,omitempty"`
}

//Stream is the event associated to an encoding instance
type Stream struct {
	ID int `json:"id,omitempty"`
}
