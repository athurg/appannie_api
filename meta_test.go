package appannie

import (
	"testing"
)

func TestCountryMeta(t *testing.T) {
	resp, err := testClient.CountryMeta()
	if err != nil {
		t.Error(err)
		return
	}

	t.Log("Countries:")
	for _, info := range resp.CountryList {
		t.Log("   ", info)
	}

	t.Log("Regions")
	for _, info := range resp.RegionList {
		t.Log("   ", info)
	}
}
