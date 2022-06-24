package danskenergi

import (
	"net/http"

	"github.com/go-resty/resty/v2"
)

const ApiEndpoint string = "https://api.elnet.danskenergi.dk/api"

type Client struct {
	*resty.Client
}

func New() *Client {
	return NewWithClient(ApiEndpoint, http.DefaultClient)
}

func NewWithClient(address string, hc *http.Client) *Client {
	rc := *resty.NewWithClient(hc)
	rc.SetBaseURL(address)
	return &Client{&rc}
}
