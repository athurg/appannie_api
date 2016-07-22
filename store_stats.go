package appannie

import (
	"fmt"
	"net/url"
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

type RatingInfo struct {
	Average     float32 `json:"average"`
	Star5Count  int     `json:"star_5_count"`
	Star4Count  int     `json:"star_4_count"`
	Star3Count  int     `json:"star_3_count"`
	Star2Count  int     `json:"star_2_count"`
	Star1Count  int     `json:"star_1_count"`
	RatingCount int     `json:"rating_count"`
}

type ProductRatingResponse struct {
	APIResponse
	PagedAPIResponse
	AppName string `json:"app_name"`
	Ratings []struct {
		Country       string     `json:"country"`
		AllRating     RatingInfo `json:"all_ratings"`
		CurrentRating RatingInfo `json:"current_ratings"` //iOS and mac Only, current version ratings
	} `json:"ratings"`
}

func (cli *Client) ProductRatings(vertical, market, asset string, productId, page int) (info ProductRatingResponse, err error) {
	q := url.Values{"page_index": []string{fmt.Sprintf("%d", page)}}

	path := fmt.Sprintf("/%s/%s/%s/%d/ratings", vertical, market, asset, productId)
	err = cli.request(path, q, &info)

	return
}

type ProductRankResponse struct {
	APIResponse
	PagedAPIResponse
	ProductName  string `json:"product_name"`
	Device       string `json:"device"`
	UpdateTime   map[string]string
	ProductRanks []struct {
		Country  string
		Category string
		Feed     string
		Interval string
		Ranks    map[string]int
	} `json:"product_ranks"`
}

//Parameter q should include params below:
//    start_date: yyyy-mm-dd
//    end_date: yyyy-mm-dd
//    interval: default daily, or hourly
//    countries: default all, refer CountryMeta()
//    category: default all, refer CategoryMeta()
//    feed: free | paid | grossing | new | top_new_free | top_new_paid | new_rising
//    device: iphone | ipad | mac | android | x86 | x64 | arm
func (cli *Client) ProductRank(vertical, market, asset string, productId int, q url.Values) (info ProductRankResponse, err error) {
	path := fmt.Sprintf("/%s/%s/%s/%d/ranks", vertical, market, asset, productId)
	err = cli.request(path, q, &info)

	return
}

type ProductFeatureResponse struct {
	APIResponse
	PagedAPIResponse
	ProductName string `json:"product_name"`
	Features    []struct {
		Level    int
		Position int
		Date     string
		Country  string
		Device   string
		Section  string
	}
}

//Parameter q should include params below:
//    start_date: yyyy-mm-dd
//    end_date: yyyy-mm-dd
//    countries: default all, refer CountryMeta()
//    page_index: 0
func (cli *Client) ProductFeature(vertical, market, asset string, productId int, q url.Values) (info ProductFeatureResponse, err error) {
	path := fmt.Sprintf("/%s/%s/%s/%d/features", vertical, market, asset, productId)
	err = cli.request(path, q, &info)

	return
}
