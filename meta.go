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
