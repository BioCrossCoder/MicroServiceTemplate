package infra

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"sync"
)

type HttpClient interface {
	Get(url string) (resBody []byte, err error)
	Post(url string, reqBody any) (resBody []byte, err error)
}

var (
	hc     HttpClient
	hcOnce sync.Once
)

type httpClient struct {
	engine *http.Client
}

func NewHttpClient() HttpClient {
	hcOnce.Do(func() {
		hc = &httpClient{engine: &http.Client{}}
	})
	return hc
}

func (c *httpClient) Get(url string) (resBody []byte, err error) {
	resp, err := c.engine.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	resBody, err = io.ReadAll(resp.Body)
	return
}

func (c *httpClient) Post(url string, reqBody any) (resBody []byte, err error) {
	data, err := json.Marshal(reqBody)
	if err != nil {
		return
	}
	resp, err := c.engine.Post(url, "application/json", bytes.NewReader(data))
	if err != nil {
		return
	}
	defer resp.Body.Close()
	resBody, err = io.ReadAll(resp.Body)
	return
}
