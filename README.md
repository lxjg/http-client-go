# http-client-go

## 初始化client
>* 无须用户验证
```python
import (
	...
	"github.com/lxjg/http-client-go/sdk"
	"github.com/lxjg/http-client-go/sdk/requests"
	...
)
client, err := sdk.NewClient()
```
>* 需要用户验证
```python
accessKeyID 为用户凭证ID，accessKeySecret 密钥

client, err := sdk.NewClientWithAccessToken(accessKeyID, accessKeySecret)
```

## 初始化请求体
```python
import (
	...
	"github.com/lxjg/http-client-go/sdk/requests"
	...
)
request := requests.NewCommonRequest()
request.Scheme = "HTTPS" // 默认为HTTP
request.Method = "POST" // 默认为GET
request.Version = "apiversion" //api 版本 
request.FormParams = map[string]string // POST 参数数组
request.QueryParams = map[string]string // GET 参数数组
```

## 启动请求
```python
res, err := client.ProcessCommonRequest(request)
res.GetHttpContentBytes() // 返回[]byte
res.GetHttpContentString() // 返回字符串
```