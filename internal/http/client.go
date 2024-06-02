package http

import "net/http"

type Client interface {
	DoRequest(url string) (int, error)
}

type httpClient struct{}

func NewHttpClient() Client {
	return &httpClient{}
}

func (c *httpClient) DoRequest(url string) (int, error) {
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	return resp.StatusCode, nil
}
