package appannie

import (
	"fmt"
)

//NOTE:
//You can only call this API for products which you have connected to Analytics or products that have been shared with you through App Annie
//Apps that are no longer published will return limited information.
func (cli *Client) ProductDetail(vertical, market, asset string, productId int) (info ProductInfo, err error) {
	var resp struct {
		APIResponse
		Product ProductInfo
	}

	path := fmt.Sprintf("/%s/%s/%s/%d/details", vertical, market, asset, productId)
	err = cli.request(path, nil, &resp)
	if err != nil {
		return
	}

	return resp.Product, err
}
