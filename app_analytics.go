package appannie

import (
	"fmt"
	"net/url"
)

type AccountInfo struct {
	AccountId      int `json:"account_id"`
	Vertical       string
	Market         string
	AccountName    string `json:"account_name"`
	PublisherName  string `json:"publisher_name"`
	FirstSalesDate string `json:"first_sales_date"`
	LastSalesDate  string `json:"last_sales_date"`
	AccountStatus  string `json:"account_status"`
}

//TODO:for large responses, need to merge all pages
func (cli *Client) Accounts() ([]AccountInfo, error) {
	var resp struct {
		APIResponse
		PagedAPIResponse
		Accounts []AccountInfo
	}

	q := url.Values{"page_index": []string{"0"}}
	err := cli.request("/accounts", q, &resp)
	if err != nil {
		return nil, err
	}

	return resp.Accounts, err
}

//TODO:for large responses, need to merge all pages
func (cli *Client) AccountProducts(accountId int) ([]ProductInfo, error) {
	var resp struct {
		APIResponse
		PagedAPIResponse
		Products []ProductInfo
	}

	path := fmt.Sprintf("/accounts/%d/products", accountId)
	q := url.Values{"page_index": []string{"0"}}
	err := cli.request(path, q, &resp)
	if err != nil {
		return nil, err
	}

	return resp.Products, err
}
