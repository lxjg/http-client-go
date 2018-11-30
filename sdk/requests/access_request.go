package requests

import (
	"io"
)

const (
	HTTP  = "HTTP"
	HTTPS = "HTTPS"

	GET     = "GET"
	PUT     = "PUT"
	POST    = "POST"
	DELETE  = "DELETE"
	HEAD    = "HEAD"
	OPTIONS = "OPTIONS"

	JSON = "application/json"
	XML  = "application/xml"
	Raw  = "application/octet-stream"
	Form = "application/x-www-form-urlencoded"

	Header = "Header"
	Query  = "Query"
	Body   = "Body"
	Path   = "Path"

	HeaderSeparator = "\n"
)

// AcsRequest interface
type AcsRequest interface {
	GetScheme() string
	GetMethod() string
	GetDomain() string
	GetHeaders() map[string]string
	GetQueryParams() map[string]string
	GetFormParams() map[string]string
	GetContent() []byte
	GetBodyReader() io.Reader
	GetVersion() string
	GetActionName() string
	GetAcceptFormat() string

	SetDomain(domain string)
	SetContent(content []byte)
	SetScheme(scheme string)
	BuildURL() string
	BuildQueries() string

	addHeaderParam(key, value string)
	addQueryParam(key, value string)
	addFormParam(key, value string)
	addPathParam(key, value string)
}

// base class
type baseRequest struct {
	Scheme       string
	Method       string
	Domain       string
	version      string
	actionName   string
	AcceptFormat string

	QueryParams map[string]string
	Headers     map[string]string
	FormParams  map[string]string
	Content     []byte

	queries string
}

func (request *baseRequest) GetQueryParams() map[string]string {
	return request.QueryParams
}

func (request *baseRequest) GetFormParams() map[string]string {
	return request.FormParams
}

func (request *baseRequest) GetContent() []byte {
	return request.Content
}

func (request *baseRequest) GetVersion() string {
	return request.version
}

func (request *baseRequest) GetActionName() string {
	return request.actionName
}

func (request *baseRequest) SetContent(content []byte) {
	request.Content = content
}

func (request *baseRequest) addHeaderParam(key, value string) {
	request.Headers[key] = value
}

func (request *baseRequest) addQueryParam(key, value string) {
	request.QueryParams[key] = value
}

func (request *baseRequest) addFormParam(key, value string) {
	request.FormParams[key] = value
}

func (request *baseRequest) GetAcceptFormat() string {
	return request.AcceptFormat
}

func (request *baseRequest) GetScheme() string {
	return request.Scheme
}

func (request *baseRequest) SetScheme(scheme string) {
	request.Scheme = scheme
}

func (request *baseRequest) GetMethod() string {
	return request.Method
}

func (request *baseRequest) GetDomain() string {
	return request.Domain
}

func (request *baseRequest) SetDomain(host string) {
	request.Domain = host
}

func (request *baseRequest) GetHeaders() map[string]string {
	return request.Headers
}

func (request *baseRequest) SetContentType(contentType string) {
	request.addHeaderParam("Content-Type", contentType)
}

func (request *baseRequest) GetContentType() (contentType string, contains bool) {
	contentType, contains = request.Headers["Content-Type"]
	return
}

func defaultBaseRequest() (request *baseRequest) {
	request = &baseRequest{
		Scheme:       HTTP,
		AcceptFormat: "JSON",
		Method:       GET,
		QueryParams:  make(map[string]string),
		Headers: map[string]string{
			"x-sdk-client":      "golang/1.0.0",
			"x-sdk-invoke-type": "normal",
			"Accept-Encoding":   "identity",
		},
		FormParams: make(map[string]string),
	}
	return
}
