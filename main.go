package httpclient

import (
	"io"
	"io/ioutil"
	"net/http"
)

type HttpClient struct {
	Header map[string]string
}

func NewHttpClient() *HttpClient {
	return &HttpClient{}
}

// SetHeader 设置请求数据类型
func (c *HttpClient) SetHeader(k string, v string) *HttpClient {
	if c.Header == nil {
		c.Header = make(map[string]string)
	}
	c.Header[k] = v

	return c
}

// SetJWTAuth 设置jwt token
func (c *HttpClient) SetJWTAuth(token string) *HttpClient {
	return c.SetHeader("Authorization", token)
}

// Get 请求
func (c *HttpClient) Get(requestLine string) ([]byte, error) {
	return c.Do("GET", requestLine, nil)
}

// Post 请求
func (c *HttpClient) Post(requestLine string, body io.Reader) ([]byte, error) {
	return c.Do("POST", requestLine, body)
}

// Patch 请求
func (c *HttpClient) Patch(requestLine string, body io.Reader) ([]byte, error) {
	return c.Do("PATCH", requestLine, body)
}

// Do 开始请求，并返回结果
func (c *HttpClient) Do(method, url string, body io.Reader) ([]byte, error) {
	req, err := http.NewRequest(method, url, body)

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
