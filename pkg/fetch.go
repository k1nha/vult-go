package main

import (
	"io"
	"net/http"
)

type FetchConfig struct {
	method string
	body   io.Reader
}

func Fetch(url string, config FetchConfig) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest(config.method, url, config.body)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
