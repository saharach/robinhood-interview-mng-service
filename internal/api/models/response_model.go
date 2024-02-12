package models

type APIResponse struct {
	Success    bool        `json:"success"`
	Message    string      `json:"message,omitempty"`
	Data       interface{} `json:"data,omitempty"`
	Errors     []string    `json:"errors,omitempty"`
	StatusCode int         `json:"statusCode"`
}
type LoginResponse struct {
	Success    bool   `json:"success"`
	Token      string `json:"token"`
	StatusCode int    `json:"statusCode"`
}

type Pagination struct {
	Total       int64 `json:"total"`
	PerPage     int   `json:"per_page"`
	CurrentPage int   `json:"current_page"`
	LastPage    int64 `json:"last_page"`
	FirstPage   int   `json:"first_page"`
}
