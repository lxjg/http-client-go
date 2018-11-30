package utils

import (
	"net/url"
)

// GetURLFormedMap 获取请求服务的参数
func GetURLFormedMap(source map[string]string) (urlEncoded string) {
	urlEncoder := url.Values{}
	for key, value := range source {
		urlEncoder.Add(key, value)
	}
	urlEncoded = urlEncoder.Encode()
	return
}
