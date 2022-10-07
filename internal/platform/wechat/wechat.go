package wechat

import (
	"github.com/go-resty/resty/v2"
)

type WechatClient struct {
	appId     string
	appSecret string
	Http      *resty.Client
}

func New(appId string, appSecret string) *WechatClient {
	HttpClient := resty.New()

	wechatClient := WechatClient{
		appId,
		appSecret,
		HttpClient,
	}

	return &wechatClient
}

type JsCodeToSessionResult struct {
	SessionKey string `json:"session_key"`
	UnionId    string `json:"unionid"`
	OpenId     string `json:"openid"`
	ErrCode    int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}

func (w *WechatClient) JsCodeToSession(code string) (*JsCodeToSessionResult, error) {
	r := JsCodeToSessionResult{}

	resp, err := w.Http.R().
		SetResult(&r).
		SetError(&r).
		ForceContentType("application/json").
		SetQueryParams(map[string]string{
			"appid":      w.appId,
			"secret":     w.appSecret,
			"js_code":    code,
			"grant_type": "authorization_code",
		}).
		Get("https://api.weixin.qq.com/sns/jscode2session")

	x := resp.Result().(*JsCodeToSessionResult)

	return x, err
}
