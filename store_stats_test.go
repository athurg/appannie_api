package appannie

import (
	"net/url"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestProductFeature(t *testing.T) {
	end := time.Now()
	start := end.AddDate(0, 0, -7)

	productId, _ := strconv.Atoi(os.Getenv("PRODUCT_ID"))

	q := url.Values{}
	q.Set("start_date", start.Format("2006-01-02"))
	q.Set("end_date", end.Format("2006-01-02"))
	q.Set("countries", "CN+RU")

	resp, err := testClient.ProductFeature("apps", "ios", "app", productId, q)
	if err != nil {
		t.Error(err)
		return
	}

	for _, f := range resp.Features {
		t.Logf("%s %s %-10s: Rank %4d by %2d Click in [%s]", f.Date, f.Country, f.Device, f.Position, f.Level, f.Section)
	}
	return
}

func TestProductReview(t *testing.T) {
	end := time.Now()
	start := end.AddDate(0, -1, 0)

	productId, _ := strconv.Atoi(os.Getenv("PRODUCT_ID"))

	q := url.Values{}
	q.Set("start_date", start.Format("2006-01-02"))
	q.Set("end_date", end.Format("2006-01-02"))

	resp, err := testClient.ProductReview("apps", "ios", "app", productId, q)
	if err != nil {
		t.Error(err)
		return
	}

	for _, r := range resp.Reviews {
		t.Logf("%-5s @v%-7s by %s", strings.Repeat("⭐️", r.Rating), r.Version, r.Reviewer)
		t.Log("[", r.Title, "]")
		t.Log(r.Text)
		t.Log("")
		t.Log("")
	}
	return
}
