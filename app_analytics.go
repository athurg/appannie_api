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

type IapInfo struct {
	Name string
	Sku  string
	Type string
}

//TODO:for large responses, need to merge all pages
func (cli *Client) AccountProductIaps(accountId, productId int) ([]IapInfo, error) {
	var resp struct {
		APIResponse
		PagedAPIResponse
		Iaps []IapInfo
	}

	path := fmt.Sprintf("/accounts/%d/products/%d/iaps", accountId, productId)
	q := url.Values{"page_index": []string{"0"}}
	err := cli.request(path, q, &resp)
	if err != nil {
		return nil, err
	}

	return resp.Iaps, err
}

type AdInfo struct {
	Name         string
	Market       string
	AdItemId     string `json:"ad_item_id"`
	AdItemType   string `json:"ad_item_type"`
	ConnectedApp []struct {
		Vertical  string
		Market    string
		ProductId int `json:"product_id"`
	} `json:"connected_app"`
}

//TODO:for large responses, need to merge all pages
//TODO:Need to be testing
//param adItemType should be site|campaign, default by both of them
func (cli *Client) AccountAdvertising(accountId int, adItemType string) ([]AdInfo, error) {
	var resp struct {
		APIResponse
		PagedAPIResponse
		AdItems []AdInfo `json:"ad_items"`
	}

	q := url.Values{"page_index": []string{"0"}}
	if adItemType != "" {
		q.Set("ad_item_type", adItemType)
	}

	path := fmt.Sprintf("/accounts/%d/ad_items", accountId)
	err := cli.request(path, q, &resp)
	if err != nil {
		return nil, err
	}

	return resp.AdItems, err
}

//TODO:for large responses, need to merge all pages
//TODO:Need to be testing
//param adItemType should be site|campaign, default by both of them
func (cli *Client) ProductAdvertising(vertical, market, asset string, productId int, adItemType string) ([]AdInfo, error) {
	var resp struct {
		APIResponse
		PagedAPIResponse
		AdItems []AdInfo `json:"ad_items"`
	}

	q := url.Values{"page_index": []string{"0"}}
	if adItemType != "" {
		q.Set("ad_item_type", adItemType)
	}

	path := fmt.Sprintf("/%s/%s/%s/%d/ad_items", vertical, market, asset, productId)
	err := cli.request(path, q, &resp)
	if err != nil {
		return nil, err
	}

	return resp.AdItems, err
}

type UserAdvertisingSalesResponse struct {
	APIResponse
	PagedAPIResponse
	Currency  string
	UserId    string `json:"user_id"`
	SalesList []struct {
		AdAccount  int `json:"ad_account"`
		Date       string
		Country    string
		Market     string
		AdItemType string `json:"ad_item_type"`
		ProductId  string `json:"product_id"`
		AdItemId   string `json:"ad_item_id"`
		Metrics    struct {
			//Revenue Metric only fields
			Revenue        float32
			Ecpm           float32
			Ecpc           float32
			FillRate       float32 `json:"fill_rate"`
			NumSites       int     `json:"num_sites"`
			Requests       int
			MatchedRequest int `json:"matched_request"`

			//Expense Metric only fields
			Installs     int
			Expense      float32
			Ctr          float32
			Cvr          float32
			Ecpi         float32
			NumCampaigns int `json:"num_campaigns"`

			//Both Revenue & Expense Metric fields
			Impressions int
			Clicks      int
			Market      string
			ProductId   int `json:"product_id"`
		}
	} `json:"sales_list"`
}

//TODO:for large responses, need to merge all pages
//TODO:Need to be testing
//param adItemType should be site|campaign, default by both of them
func (cli *Client) UserAdvertisingSales(vertical string, q url.Values) (info UserAdvertisingSalesResponse, err error) {
	err = cli.request("/"+vertical+"/sales", q, &info)
	return
}
