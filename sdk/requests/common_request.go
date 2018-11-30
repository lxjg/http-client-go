package requests

import (
	"io"
	"strings"

	"github.com/lxjg/http-client-go/sdk/utils"
)

// CommonRequest 通用http request对象
type CommonRequest struct {
	*baseRequest

	Version    string
	ActionName string
}

// NewCommonRequest 初始化CommonRequest
func NewCommonRequest() (request *CommonRequest) {
	request = &CommonRequest{
		baseRequest: defaultBaseRequest(),
	}
	request.Headers["x-sdk-invoke-type"] = "common"
	return
}

// BuildURL 建立请求路由
func (request *CommonRequest) BuildURL() string {
	return strings.ToLower(request.Scheme) + "://" + request.Domain + "/" + request.Version + "/" + request.ActionName + request.BuildQueries()
}

// GetBodyReader 获取request body 参数
func (request *CommonRequest) GetBodyReader() io.Reader {
	if request.FormParams != nil && len(request.FormParams) > 0 {
		formString := utils.GetURLFormedMap(request.FormParams)
		return strings.NewReader(formString)
	}

	return strings.NewReader("")
}

// BuildQueries 获取路由参数
func (request *CommonRequest) BuildQueries() string {
	request.queries = "?" + utils.GetURLFormedMap(request.QueryParams)
	return request.queries
}
