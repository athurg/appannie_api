package appannie

import (
	"testing"
)

func TestSharingProducts(t *testing.T) {
	//测试前请输入有效的AppAnnieKey
	client := New("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx", "")
	sharings, err := client.SharingProducts()
	if err != nil {
		t.Error(err)
		return
	}

	for _, s := range sharings {
		t.Logf("分享者:%s(%d) 类型 %s\n", s.OwnerName, s.OwnerAccountId, s.Vertical)
		for _, p := range s.Products {
			prefix := "❌"
			if p.Status {
				prefix = "✅"
			}
			t.Logf("\t %s %-14d %s, 渠道%s\n", prefix, p.ProductId, p.ProductName, p.Market)
		}
		t.Logf("\n")
	}
}
