package appannie

import (
	"net/url"
)

type ProductInfo struct {
	ProductId      int    `json:"product_id"`
	ProductName    string `json:"product_name"`
	Icon           string `json:"icon"`
	Market         string `json:"market"`
	Status         bool   `json:"status"`
	FirstSalesDate string `json:"first_sales_date"`
	LastSalesDate  string `json:"last_sales_date"`

	Vertical                 string   `json:"vertical"`
	BundleId                 string   `json:"bundle_id"`
	PublisherName            string   `json:"publisher_name"`
	Description              string   `json:"description"`
	CurrentVersion           string   `json:"current_version"`
	ReleaseDate              string   `json:"release_date"`
	LastUpdate               string   `json:"last_update"`
	Size                     string   `json:"size"`
	Languages                string   `json:"languages"`
	MainCategory             string   `json:"main_category"`
	OtherCategories          string   `json:"other_categories"`
	PublisherId              int      `json:"publisher_id"`
	Price                    float32  `json:"price"`
	PurchasedSeparatelyPrice float32  `json:"purchased_separately_price"` //iOS only,list
	Unpublished              bool     `json:"unpublished"`
	HasIap                   bool     `json:"has_iap"`
	OriginalIcon             string   `json:"original_icon"`            //icon url from app store
	ProductType              string   `json:"product_type"`             //iOS only App|Bundle
	Seller                   string   `json:"seller"`                   //iOS only
	FamilySharing            string   `json:"family_sharing"`           //iOS only, Yes|No
	AppsInThisBundle         []string `json:"apps_in_this_bundle"`      //iOS only
	BundlesContainThisApp    []string `json:"bundles_contain_this_app"` //iOS only
	SupportedDeviceList      []string `json:"supported_device_list"`    //iOS only
}

type SharingInfo struct {
	Vertical       string `json:"vertical"`
	OwnerAccountId int    `json:"owner_account_id"`
	OwnerName      string `json:"owner_name"`
	Products       []ProductInfo
}

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
