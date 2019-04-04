package kickbox

import (
	"net/http"
	"time"
)

// Client ...
type Client struct {
	apiKey string
	client *http.Client
}

// Result ...
type Result struct {
	Result     string  `json:"result"`
	Reason     string  `json:"reason"`
	Role       bool    `json:"role"`
	Disposable bool    `json:"disposable"`
	AcceptAll  bool    `json:"accept_all"`
	Free       bool    `json:"free"`
	DidYouMean string  `json:"did_you_mean"`
	Sendex     float32 `json:"sendex"`
	Email      string  `json:"email"`
	User       string  `json:"user"`
	Domain     string  `json:"domain"`
	Success    bool    `json:"success"`
	Message    string  `json:"message"`
}

// Error ....
type Error struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Code    int
	err     error
}

// New ...
func New(apiKey string) *Client {
	client := &Client{apiKey: apiKey, client: &http.Client{
		Timeout: time.Second * 60,
	}}

	return client
}
