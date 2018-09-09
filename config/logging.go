package config

type logging struct {
	LogRequests  bool `json:"logRequests"`
	LogResponses bool `json:"logResponses"`
}
