package appannie

import (
	"testing"
	"time"
)

func TestProductSales(t *testing.T) {
	//测试前请输入有效的AppAnnieKey和账户、产品ID
	client := New("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx", "")
	var accountId int64 = 11111
	var productId int64 = 22222222222222
	end := time.Now()
	start := end.AddDate(0, -1, 0)

	t.Log("从", start, "到", end)

	resp, err := client.ProductSales(accountId, productId, start, end)
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("销售记录%d条\n", len(resp.SalesList))
	for _, rec := range resp.SalesList {
		t.Logf("%+v", rec)
	}

	t.Log("DONE")
}
