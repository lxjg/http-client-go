package responses

import (
	"errors"
	"io/ioutil"
	"net/http"
)

// AcsResponse response接口约定
type AcsResponse interface {
	IsSuccess() bool
	GetHttpStatus() int
	GetHttpContentString() string
	GetHttpContentBytes() []byte
	parseFromHttpResponse(httpResponse *http.Response) error
}

// Unmarshal object from http response body to target Response
func Unmarshal(response AcsResponse, httpResponse *http.Response) (err error) {
	err = response.parseFromHttpResponse(httpResponse)
	if err != nil {
		return
	}
	if !response.IsSuccess() {
		err = errors.New(response.GetHttpContentString())
		return
	}

	return
}

// BaseResponse 基础response结构体
type BaseResponse struct {
	httpStatus         int
	httpContentString  string
	httpContentBytes   []byte
	originHTTPResponse *http.Response
}

// CommonResponse 通用response结构体
type CommonResponse struct {
	*BaseResponse
}

// NewCommonResponse 初始化CommonResponse
func NewCommonResponse() (response *CommonResponse) {
	return &CommonResponse{
		BaseResponse: &BaseResponse{},
	}
}

func (baseResponse *BaseResponse) GetHttpStatus() int {
	return baseResponse.httpStatus
}

func (baseResponse *BaseResponse) IsSuccess() bool {
	if baseResponse.GetHttpStatus() >= 200 && baseResponse.GetHttpStatus() < 300 {
		return true
	}

	return false
}

func (baseResponse *BaseResponse) GetHttpContentString() string {
	return baseResponse.httpContentString
}

func (baseResponse *BaseResponse) GetHttpContentBytes() []byte {
	return baseResponse.httpContentBytes
}

func (baseResponse *BaseResponse) parseFromHttpResponse(httpResponse *http.Response) (err error) {
	defer httpResponse.Body.Close()
	body, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return
	}
	baseResponse.httpStatus = httpResponse.StatusCode
	baseResponse.httpContentBytes = body
	baseResponse.httpContentString = string(body)
	baseResponse.originHTTPResponse = httpResponse
	return
}
