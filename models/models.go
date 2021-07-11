package models

// Endpoint is single route for API
type Endpoint struct {
	Route    string `json:"route"`
	JSONPath string `json:"jsonpath"`
}

// Config is model for all endpoints
type Config struct {
	Endpoints []Endpoint `json:"endpoints"`
}
