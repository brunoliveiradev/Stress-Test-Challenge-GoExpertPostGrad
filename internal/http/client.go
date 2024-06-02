package http

import (
	"net/http"
)

type Client interface {
	DoRequest(url string) (int, error)
}

type httpClient struct {
	client *http.Client
}

func NewHttpClient() Client {
	return &httpClient{
		client: &http.Client{},
	}
}

func (c *httpClient) DoRequest(url string) (int, error) {
	resp, err := c.client.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	return resp.StatusCode, nil
}
