package httpclient

import (
	"io/ioutil"
	"net/http"
)

type HttpClient struct {
	Header map[string]string
}

func NewHttpClient() *HttpClient {
	return &HttpClient{}
}

func (c *HttpClient) SetHeader(k string, v string) *HttpClient {
	if c.Header == nil {
		c.Header = make(map[string]string)
	}
	c.Header[k] = v

	return c
}

func (c *HttpClient) SetJWTAuth(token string) *HttpClient {
	return c.SetHeader("Authorization", token)
}

// Get 请求
func (c *HttpClient) Get(requestLine string) ([]byte, error) {
	req, err := http.NewRequest("GET", requestLine, nil)
	if err != nil {
		return nil, err
	}

	if c.Header != nil {
		for k, v := range c.Header {
			req.Header.Set(k, v)
		}
	}

	client := http.Client{}
	response, err := client.Do(req)
	if err != nil || response.StatusCode != http.StatusOK {
		return nil, err
	}

	defer response.Body.Close()
	return ioutil.ReadAll(response.Body)
}
