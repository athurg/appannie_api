package appannie

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
	if info.Code != 200 {
		return
	}

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
	if info.Code != 200 {
		return
	}

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
	if info.Code != 200 {
		return
	}

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
	if info.Code != 200 {
		return
	}

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
	if info.Code != 200 {
		return
	}

	return
}
