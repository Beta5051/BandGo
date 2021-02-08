package BandGo

import "net/http"

type Client struct {
	Token string
	Debug bool
	http  *http.Client
}

func New(token string, debug bool) *Client {
	return &Client{
		Token: token,
		Debug: debug,
		http:  &http.Client{},
	}
}