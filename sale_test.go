package appannie

import (
	"os"
	"strconv"
	"testing"
	"time"
)

func TestProductSales(t *testing.T) {
	end := time.Now()
	start := end.AddDate(0, -1, 0)

	var accountId, productId int
	accountId, _ = strconv.Atoi(os.Getenv("ACCOUNT_ID"))
	productId, _ = strconv.Atoi(os.Getenv("PRODUCT_ID"))

	resp, err := testClient.ProductSales(accountId, productId, start, end)
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("Sale records x%d , between %s and %s.\n", len(resp.SalesList), start, end)
	for _, rec := range resp.SalesList {
		t.Logf("%+v", rec)
	}

	t.Log("DONE")
}
