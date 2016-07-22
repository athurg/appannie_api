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

func (cli *Client) CountryMeta() ([]CountryInfo, []RegionInfo, error) {
	var resp struct {
		APIResponse
		CountryList []CountryInfo `json:"country_list"`
		RegionList  []RegionInfo  `json:"region_list"`
	}

	err := cli.request("/meta/countries", nil, &resp)
	if err != nil {
		return nil, nil, err
	}

	return resp.CountryList, resp.RegionList, nil
}

//Available vertical: apps
//Available market: ios | mac | google-play | amazon-appstore | windows-phone | windows-store
func (cli *Client) CategoryMeta(vertical, market string) ([]string, []string, error) {
	var resp struct {
		APIResponse
		Categories         []string
		AppAnnieCategories []string `json:"appannie_categories"`
	}

	err := cli.request("/meta/"+vertical+"/"+market+"/categories", nil, &resp)
	if err != nil {
		return nil, nil, err
	}

	return resp.Categories, resp.AppAnnieCategories, nil
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

func (cli *Client) CurrencyMeta() ([]CurrencyInfo, error) {
	var resp struct {
		APIResponse
		CurrencyList []CurrencyInfo `json:"currency_list"`
	}

	err := cli.request("/meta/currencies", nil, &resp)
	if err != nil {
		return nil, err
	}

	return resp.CurrencyList, nil
}

type DeviceInfo struct {
	DeviceCode string `json:"device_code"`
	DeviceName string `json:"device_name"`
}

//Available vertical: apps
//Available market: ios | mac | google-play | windows-store
func (cli *Client) DeviceMeta(vertical, market string) ([]DeviceInfo, error) {
	var resp struct {
		APIResponse
		Devices []DeviceInfo `json:"devices"`
	}

	err := cli.request("/meta/"+vertical+"/"+market+"/devices", nil, &resp)
	if err != nil {
		return nil, err
	}

	return resp.Devices, nil
}

type FeedInfo struct {
	FeedCode string `json:"feed_code"`
	FeedName string `json:"feed_name"`
}

//Available vertical: apps
//Available market: ios | mac | google-play | amazon-appstore |  windows-phone | windows-store
func (cli *Client) FeedMeta(vertical, market string) ([]FeedInfo, error) {
	var resp struct {
		APIResponse
		Feeds []FeedInfo `json:"feeds"`
	}

	err := cli.request("/meta/"+vertical+"/"+market+"/feeds", nil, &resp)
	if err != nil {
		return nil, err
	}
	return resp.Feeds, nil
}

//Available vertical: apps
//Available market: ios | google-play
//package_codes: BundleID for iOS, class for GP, split by comma
func (cli *Client) PackageCodesToProductIds(vertical, market, package_codes string) (map[string]int, error) {
	q := url.Values{
		"package_codes": []string{package_codes},
	}

	var resp struct {
		APIResponse
		Items []struct {
			ProductId   int    `json:"product_id"`
			PackageCode string `json:"package_code"`
		}
	}

	err := cli.request("/"+vertical+"/"+market+"/package-codes2ids", q, &resp)
	if err != nil {
		return nil, err
	}

	result := make(map[string]int)
	for _, info := range resp.Items {
		result[info.PackageCode] = info.ProductId
	}

	return result, nil
}
