package strategy

import (
	"os"
	"testing"
)

func TestGenerateUrl(t *testing.T) {
	apiKey := "key"
	os.Setenv("APIKEY", apiKey)

	textSearch := TextSearch{}
	queries := []string{"tokyo", "restaurant"}
	url := textSearch.BaseURL() + textSearch.Formatquery(queries) + textSearch.Options() + "&key=" + apiKey

	expected := "https://maps.googleapis.com/maps/api/place/textsearch/json?query=tokyo+restaurant&language=ja&key=key"

	if url != expected {
		t.Errorf("expected: %s\ngot: %s", expected, url)
	}
}
