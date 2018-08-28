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
	requestParams  map[string]string
	requestHeaders map[string]string
}

func (rc *requestConfig) MakeGet(url string) (string, error) {
	resty.SetDebug(true)
	resty.SetHeaders(rc.requestHeaders)
	resty.SetQueryParams(rc.requestParams)

	resp, err := resty.R().Get(rc.baseUrl + url)
	if err != nil {
		return "nil", err
	}
	return string(resp.Body()), nil
}

func (rc *requestConfig) MakeDelete(url string) (string, error) {
	resty.SetDebug(true)
	resty.SetHeaders(rc.requestHeaders)

	resp, err := resty.R().Delete(rc.baseUrl + url)
	if err != nil {
		return "nil", err
	}
	return string(resp.Body()), nil
}

func (rc *requestConfig) MakePost(url string) (string, error) {
	resty.SetDebug(true)
	resty.SetHeaders(rc.requestHeaders)
	resty.SetFormData(rc.requestParams)

	resp, err := resty.R().Post(rc.baseUrl + url)

	if err != nil {
		return "nil", err
	}
	return string(resp.Body()), nil
}

func (rc *requestConfig) MakePut(url string) (string, error) {
	resty.SetDebug(true)
	resty.SetHeaders(rc.requestHeaders)
	resty.SetFormData(rc.requestParams)

	resp, err := resty.R().Put(rc.baseUrl + url)

	if err != nil {
		return "nil", err
	}
	return string(resp.Body()), nil
}

func NewRequestUtility(baseUrl string, requestMethod string, requestParams map[string]string, requestHeaders map[string]string) RequestUtility {
	return &requestConfig{
		baseUrl, requestMethod, requestParams, requestHeaders,
	}

}
