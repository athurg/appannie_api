package appannie

import (
	"testing"
)

func TestCountryMeta(t *testing.T) {
	countries, regions, err := testClient.CountryMeta()
	if err != nil {
		t.Error(err)
		return
	}

	t.Log("Countries:")
	for _, info := range countries {
		t.Log("   ", info)
	}

	t.Log("Regions")
	for _, info := range regions {
		t.Log("   ", info)
	}
}

func TestCategoryMeta(t *testing.T) {
	categories, appAnnieCategories, err := testClient.CategoryMeta("apps", "google-play")
	if err != nil {
		t.Error(err)
		return
	}

	t.Log("Categories:")
	for _, info := range categories {
		t.Log("   ", info)
	}

	t.Log("AppAnnie Categories:")
	for _, info := range appAnnieCategories {
		t.Log("   ", info)
	}
}

func TestCurrencyMeta(t *testing.T) {
	metas, err := testClient.CurrencyMeta()
	if err != nil {
		t.Error(err)
		return
	}

	t.Log("Categories:")
	for _, m := range metas {
		t.Log("   ", m)
	}
}
