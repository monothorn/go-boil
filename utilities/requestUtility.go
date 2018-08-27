package utilities

import (
	"gopkg.in/resty.v1"
)

type RequestUtility interface {
	MakeGet(url string) (string, error)
}

type requestConfig struct {
	baseUrl        string
	requestMethod  string
	requestParams  string
	requestHeaders map[string]string
}

func (rc *requestConfig) MakeGet(url string) (string, error) {
	resp, err := resty.R().Get("http://httpbin.org/get")
	if err != nil {
		return "nil", err
	}
	return string(resp.Body()), nil
}

func NewRequestUtility(baseUrl string, requestMethod string, requestParams string, requestHeaders map[string]string) RequestUtility {
	return &requestConfig{
		baseUrl, requestMethod, requestParams, requestHeaders,
	}

}
