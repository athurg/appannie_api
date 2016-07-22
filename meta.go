package appannie

import (
	"net/url"
)

type CountryInfo struct {
	CountryCode string `json:"country_code"`
	CountryName string `json:"country_name"`
}

type RegionInfo struct {
	RegionCode string `json:"region_code"`
	RegionName string `json:"region_name"`
}

type CountryMetaResponse struct {
	APIResponse
	CountryList []CountryInfo `json:"country_list"`
	RegionList  []RegionInfo  `json:"region_list"`
}

func (cli *Client) CountryMeta() (info CountryMetaResponse, err error) {
	err = cli.request("/meta/countries", nil, &info)
	return
}

type CategoryMetaResponse struct {
	APIResponse
	Categories         []string `json:"categories"`
	AppAnnieCategories []string `json:"appannie_categories"`
}

//Available vertical: apps
//Available market: ios | mac | google-play | amazon-appstore | windows-phone | windows-store
func (cli *Client) CategoryMeta(vertical, market string) (info CategoryMetaResponse, err error) {
	err = cli.request("/meta/"+vertical+"/"+market+"/categories", nil, &info)
	return
}

type MarketInfo struct {
	MarketName string `json:"market_name"`
	MarketCode string `json:"market_code"`
}

type MarketMetaResponse struct {
	APIResponse
	Verticals []struct {
		VerticalName string       `json:"vertical_name"`
		Markets      []MarketInfo `json:"markets"`
	}
}

func (cli *Client) MarketMeta() (info MarketMetaResponse, err error) {
	err = cli.request("/meta/markets", nil, &info)
	return
}

type CurrencyInfo struct {
	CurrencyCode string `json:"currency_code"`
	FullName     string `json:"full_name"`
	Symbol       string `json:"symbol"`
}

type CurrencyMetaResponse struct {
	APIResponse
	CurrencyList []CurrencyInfo `json:"currency_list"`
}

func (cli *Client) CurrencyMeta() (info CurrencyMetaResponse, err error) {
	err = cli.request("/meta/currencies", nil, &info)
	return
}

type DeviceInfo struct {
	DeviceCode string `json:"device_code"`
	DeviceName string `json:"device_name"`
}

type DeviceMetaResponse struct {
	APIResponse
	Devices []DeviceInfo `json:"devices"`
}

//Available vertical: apps
//Available market: ios | mac | google-play | windows-store
func (cli *Client) DeviceMeta(vertical, market string) (info DeviceMetaResponse, err error) {
	err = cli.request("/meta/"+vertical+"/"+market+"/devices", nil, &info)
	return
}

type FeedInfo struct {
	FeedCode string `json:"feed_code"`
	FeedName string `json:"feed_name"`
}

type FeedMetaResponse struct {
	APIResponse
	Feeds []FeedInfo `json:"feeds"`
}

//Available vertical: apps
//Available market: ios | mac | google-play | amazon-appstore |  windows-phone | windows-store
func (cli *Client) FeedMeta(vertical, market string) (info FeedMetaResponse, err error) {
	err = cli.request("/meta/"+vertical+"/"+market+"/feeds", nil, &info)
	return
}

//Available vertical: apps
//Available market: ios | google-play
//package_codes: BundleID for iOS, class for GP, split by comma
func (cli *Client) PackageCodesToProductIds(vertical, market, package_codes string) (result map[string]int, err error) {
	q := url.Values{"package_codes": []string{package_codes}}
	var resp struct {
		APIResponse
		Items []struct {
			ProductId   int    `json:"product_id"`
			PackageCode string `json:"package_code"`
		}
	}
	err = cli.request("/"+vertical+"/"+market+"/package-codes2ids", q, &resp)
	if err != nil {
		return
	}

	result = make(map[string]int)
	for _, info := range resp.Items {
		result[info.PackageCode] = info.ProductId
	}

	return
}
