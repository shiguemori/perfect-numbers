package models

import "time"

type PerfectNumberResponse struct {
	PerfectNumbers []int     `json:"perfect_numbers"`
	Count          int       `json:"count"`
	Range          string    `json:"range"`
	ProcessingTime string    `json:"processing_time"`
	Timestamp      time.Time `json:"timestamp"`
}

type ErrorResponse struct {
	Error     string    `json:"error"`
	Code      string    `json:"code,omitempty"`
	Timestamp time.Time `json:"timestamp"`
}

type HealthResponse struct {
	Status    string    `json:"status"`
	Message   string    `json:"message"`
	Version   string    `json:"version"`
	Timestamp time.Time `json:"timestamp"`
	Uptime    string    `json:"uptime"`
}

type APIInfoResponse struct {
	Name        string            `json:"name"`
	Version     string            `json:"version"`
	Description string            `json:"description"`
	Endpoints   map[string]string `json:"endpoints"`
	Timestamp   time.Time         `json:"timestamp"`
}
