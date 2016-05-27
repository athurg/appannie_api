package appannie

import (
	"encoding/json"
	"net/http"
	"net/url"
)

type Client struct {
	AppUrl      string
	AppAnnieKey string
	ApiVersion  string
}

//创建一个AppAnnie数据访问的客户端
func New(key string, version string) *Client {
	if version == "" {
		version = "v1.2"
	}

	return &Client{
		AppAnnieKey: key,
		ApiVersion:  version,
	}
}

//获取含版本号的API访问地址
func (cli *Client) ApiUrl() string {
	return "https://api.appannie.com/" + cli.ApiVersion
}

//执行API请求（自动处理Token验证）
func (cli *Client) request(path string, q url.Values, result interface{}) error {
	client := &http.Client{}

	uri := cli.ApiUrl() + path + "?" + q.Encode()
	req, err := http.NewRequest("GET", uri, nil)
	req.Header.Add("Authorization", "bearer "+cli.AppAnnieKey)

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(&result)
}
