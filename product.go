package appannie

import (
	"net/url"
)

//分享的产品信息
type ProductInfo struct {
	ProductId      int    `json:"product_id"`
	ProductName    string `json:"product_name"`
	Icon           string `json:"icon"`
	Market         string `json:"market"`
	Status         bool   `json:"status"`
	FirstSalesDate string `json:"first_sales_date"`
	LastSalesDate  string `json:"last_sales_date"`
}

//分享信息
type SharingInfo struct {
	Vertical       string `json:"vertical"`
	OwnerAccountId int    `json:"owner_account_id"`
	OwnerName      string `json:"owner_name"`
	Products       []ProductInfo
}

//分享响应信息
type SharingProductsResponse struct {
	APIResponse
	Sharings  []SharingInfo
	PageNum   int `json:"page_num"`
	PageIndex int `json:"page_index"`
	PrevPage  int `json:"prev_page"`
	NextPage  int `json:"next_page"`
}

func (cli *Client) SharingProducts() ([]SharingInfo, error) {
	q := url.Values{}
	q.Set("page_index", "0")

	var info SharingProductsResponse
	err := cli.request("/sharing/products", q, &info)

	return info.Sharings, err
}
