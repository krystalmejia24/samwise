package db

type Encoder struct {
	IP     string  `json:"ip"`
	Config *Config `json:"config,omitempty"`
	Stream *Stream `json:"event,omitempty"`
}

type Config struct {
	Username string `json:"username,omitempty"`
	APIKey   string `json:"api_key,omitempty"`
}

type Stream struct {
	ID int `json:"id,omitempty"`
}
