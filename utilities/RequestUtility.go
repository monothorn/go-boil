package thirdparty

import (
	"gopkg.in/resty.v1"
)

type RequestUtility interface {
	makeRequest(url string, data map[string]string) (string, error)
	makeJSONRequest(url string, data string) (string, error)
}

type requestConfig struct {
	baseUrl        string
	requestMethod  string
	requestHeaders map[string]string
}

func (rc *requestConfig) makeRequest(url string, data map[string]string) (string, error) {
	resty.SetDebug(true)
	resty.SetHeaders(rc.requestHeaders)
	resty.SetHostURL(rc.baseUrl)
	if rc.requestMethod == "GET" || rc.requestMethod == "DELETE" {
		resty.SetQueryParams(data)
	}
	resty.SetContentLength(true)

	switch rc.requestMethod {
	case "GET":
		resp, err := resty.R().Get(url)
		if err != nil {
			return "nil", err
		}
		return string(resp.Body()), nil
	case "POST":
		resp, err := resty.R().SetBody(data).Post(url)
		if err != nil {
			return "nil", err
		}
		return string(resp.Body()), nil
	case "PUT":
		resp, err := resty.R().SetBody(data).Put(url)
		if err != nil {
			return "nil", err
		}
		return string(resp.Body()), nil
	case "PATCH":
		resp, err := resty.R().SetBody(data).Patch(url)
		if err != nil {
			return "nil", err
		}
		return string(resp.Body()), nil
	case "DELETE":
		resp, err := resty.R().Delete(url)
		if err != nil {
			return "nil", err
		}
		return string(resp.Body()), nil
	}

	return "nil", nil
}

func (rc *requestConfig) makeJSONRequest(url string, data string) (string, error) {
	resty.SetDebug(true)
	resty.SetHeaders(rc.requestHeaders)
	resty.SetHostURL(rc.baseUrl)
	resty.SetContentLength(true)

	switch rc.requestMethod {
	case "GET":
		resp, err := resty.R().SetQueryString(data).Get(url)
		if err != nil {
			return "nil", err
		}
		return string(resp.Body()), nil
	case "POST":
		resp, err := resty.R().SetBody(data).Post(url)
		if err != nil {
			return "nil", err
		}
		return string(resp.Body()), nil
	case "PUT":
		resp, err := resty.R().SetBody(data).Put(url)
		if err != nil {
			return "nil", err
		}
		return string(resp.Body()), nil
	case "PATCH":
		resp, err := resty.R().SetBody(data).Patch(url)
		if err != nil {
			return "nil", err
		}
		return string(resp.Body()), nil
	case "DELETE":
		resp, err := resty.R().SetQueryString(data).Delete(url)
		if err != nil {
			return "nil", err
		}
		return string(resp.Body()), nil
	}

	return "nil", nil
}

func NewRequestUtility(baseUrl string, requestMethod string, requestHeaders map[string]string) RequestUtility {
	return &requestConfig{
		baseUrl, requestMethod, requestHeaders,
	}

}
