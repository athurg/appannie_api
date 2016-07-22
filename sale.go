package appannie

import (
	"fmt"
	"net/url"
	"time"
)

type SaleInfo struct {
	Date      string
	ProductId string `json:"product_id,omitempty"`
	Country   string
	Units     struct {
		Product struct{ Downloads, Updates, Refunds, Promotions int64 }
		Iap     struct{ Sales, Refunds, Promotions int }
	}
	Revenue struct {
		Product struct{ Downloads, Updates, Refunds, Promotions string }
		Iap     struct{ Sales, Refunds, Promotions string }
		Ad      string
	}
}

type IapSaleInfo struct {
	Date    string
	Country string
	Iap     string
	Units   struct{ Sales, Refunds, Promotions int }
	Revenue struct {
		Sales      string
		Refunds    string
		Promotions string
	}
}

type ProductSaleResponse struct {
	APIResponse
	Currency  string
	Vertical  string
	Market    string
	SalesList []SaleInfo    `json:"sales_list,omitempty"`
	IapSales  []IapSaleInfo `json:"iap_sales,omitempty"`

	PageNum   int `json:"page_num"`
	PageIndex int `json:"page_index"`
	PrevPage  int `json:"prev_page"`
	NextPage  int `json:"next_page"`
}

func (cli *Client) ProductSales(accountId, productId int, start, end time.Time) (info ProductSaleResponse, err error) {
	q := url.Values{}
	q.Set("break_down", "date")
	q.Set("start_date", start.Format("2006-01-02"))
	q.Set("end_date", end.Format("2006-01-02"))
	//q.Set("page_index", "0")
	//q.Set("currency", "USD")
	//q.Set("countries", "all")

	path := fmt.Sprintf("/accounts/%d/products/%d/sales", accountId, productId)
	err = cli.request(path, q, &info)

	return
}

//TODO:Need to be testing
func (cli *Client) AccountSales(accountId int, start, end time.Time) (info ProductSaleResponse, err error) {
	q := url.Values{}
	q.Set("break_down", "date")
	q.Set("start_date", start.Format("2006-01-02"))
	q.Set("end_date", end.Format("2006-01-02"))
	//q.Set("page_index", "0")
	//q.Set("currency", "USD")
	//q.Set("countries", "all")

	path := fmt.Sprintf("/accounts/%d/sales", accountId)
	err = cli.request(path, q, &info)

	return
}
