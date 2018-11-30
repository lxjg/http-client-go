package sdk

import (
	"fmt"
	"net/http"

	"github.com/lxjg/http-client-go/sdk/auth"
	"github.com/lxjg/http-client-go/sdk/requests"
	"github.com/lxjg/http-client-go/sdk/responses"
)

// Client http client
type Client struct {
	token      string
	httpClient *http.Client
}

// NewClient 初始化http client
func NewClient() (client *Client, err error) {
	client = &Client{
		httpClient: &http.Client{},
	}
	return client, nil
}

// NewClientWithAccessToken 初始化需要验证的http client
func NewClientWithAccessToken(accessKeyID, accessKeySecret string) (client *Client, err error) {
	client = &Client{}
	credential := &auth.Credential{
		AccessKeyID:     accessKeyID,
		AccessKeySecret: accessKeySecret,
	}

	err = client.InitWithOptions(credential)
	return client, err
}

// InitWithOptions 初始化Client
func (client *Client) InitWithOptions(credential *auth.Credential) (err error) {
	if err != nil {
		return err
	}

	client.httpClient = &http.Client{}
	client.token, err = auth.GenerateJWT(credential)
	return err
}

// ProcessCommonRequest 启动http请求
func (client *Client) ProcessCommonRequest(request *requests.CommonRequest) (response *responses.CommonResponse, err error) {
	response = responses.NewCommonResponse()
	err = client.DoAction(request, response)
	return
}

// DoAction 发起http请求
func (client *Client) DoAction(request *requests.CommonRequest, response responses.AcsResponse) (err error) {
	httpRequest, err := client.buildHTTPRequest(request)
	var httpResponse *http.Response
	httpResponse, err = client.httpClient.Do(httpRequest)
	err = responses.Unmarshal(response, httpResponse)
	return err
}

func (client *Client) buildHTTPRequest(request *requests.CommonRequest) (httpRequest *http.Request, err error) {
	requestMethod := request.GetMethod()
	requestURL := request.BuildURL()
	body := request.GetBodyReader()
	httpRequest, err = http.NewRequest(requestMethod, requestURL, body)
	if err != nil {
		return
	}
	for key, value := range request.GetHeaders() {
		httpRequest.Header[key] = []string{value}
	}

	if len(client.token) > 0 {
		authorization := fmt.Sprintf("Bearer %s", client.token)
		httpRequest.Header["Authorization"] = []string{authorization}
	}

	return
}
