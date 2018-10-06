package httpclient

import (
	"io/ioutil"
	"net/http"
)

// Get 请求
func Get(requestLine string) ([]byte, error) {
	response, err := http.Get(requestLine)
	if err != nil || response.StatusCode != http.StatusOK {
		return nil, err
	}

	defer response.Body.Close()
	return ioutil.ReadAll(response.Body)
}
