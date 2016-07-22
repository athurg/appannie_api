package appannie

import (
	"testing"
	"time"
)

func TestSharingProducts(t *testing.T) {
	end := time.Now()
	start := end.AddDate(0, -3, 0)

	//测试前请输入有效的AppAnnieKey
	sharings, err := testClient.SharingProducts()
	if err != nil {
		t.Error(err)
		return
	}

	for _, s := range sharings {
		for _, p := range s.Products {
			var count = -1
			resp, err := testClient.ProductSales(s.OwnerAccountId, p.ProductId, start, end)
			if err != nil {
				t.Error(err)
			} else {
				count = len(resp.SalesList)
			}

			var prefix string
			if p.Status == false {
				prefix = "❌"
			} else if count > 0 {
				prefix = "✅"
			} else {
				prefix = "⚠️"
			}
			t.Logf("\t%s  %7d-%-14d %s, 👉 %s x %d\n", prefix, s.OwnerAccountId, p.ProductId, p.ProductName, p.Market, count)
		}
		t.Logf("\n")
	}
}
