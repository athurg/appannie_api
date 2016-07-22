package appannie

import (
	"encoding/json"
	"fmt"
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

type APIResponser interface {
	Error() error
}

type APIResponse struct {
	Code         int    `json:"code"`
	ErrorMessage string `json:"error"`
}

type PagedAPIResponse struct {
	PageNum   int    `json:"page_num"`
	PageIndex int    `json:"page_index"`
	PrevPage  string `json:"prev_page"`
	NextPage  string `json:"next_page"`
}

func (resp *APIResponse) Error() error {
	if resp.Code == 200 {
		return nil
	}
	return fmt.Errorf("[%d]%s", resp.Code, resp.ErrorMessage)
}

//执行API请求（自动处理Token验证）
//TODO:目前API有30次每分钟、1000次每天的请求限制
func (cli *Client) request(path string, q url.Values, result APIResponser) error {
	client := &http.Client{}

	uri := cli.ApiUrl() + path + "?" + q.Encode()
	req, err := http.NewRequest("GET", uri, nil)
	req.Header.Add("Authorization", "bearer "+cli.AppAnnieKey)

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return err
	}

	return result.Error()
}
