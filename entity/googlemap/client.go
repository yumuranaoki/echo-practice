package googlemap

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	Entity "github.com/yumuranaoki/echo-practice/entity"
)

type APIClient struct {
	strategy
}

type strategy interface {
	formatquery([]string) string
	baseURL() string
	options() string
}

type resp struct {
	candidates []json.RawMessage
	debugLog   json.RawMessage
	status     string
}

func (client APIClient) get(keywords []string) string {
	url := client.baseURL() + client.formatquery([]string{"東京", "観光"}) + client.options()
	resp, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	myJson := `
		{
			"candidates" : [
				{
					"formatted_address" : "140 George St, The Rocks NSW 2000, Australia",
					"geometry" : {
						"location" : {
								"lat" : -33.8599358,
								"lng" : 151.2090295
						},
						"viewport" : {
								"northeast" : {
									"lat" : -33.85824767010727,
									"lng" : 151.2102470798928
								},
								"southwest" : {
									"lat" : -33.86094732989272,
									"lng" : 151.2075474201073
								}
						}
					},
					"name" : "Museum of Contemporary Art Australia",
					"opening_hours" : {
						"open_now" : false,
						"weekday_text" : []
					},
					"photos" : [
						{
								"height" : 2268,
								"html_attributions" : [
									"\u003ca href=\"https://maps.google.com/maps/contrib/113202928073475129698/photos\"\u003eEmily Zimny\u003c/a\u003e"
								],
								"photo_reference" : "CmRaAAAAfxSORBfVmhZcERd-9eC5X1x1pKQgbmunjoYdGp4dYADIqC0AXVBCyeDNTHSL6NaG7-UiaqZ8b3BI4qZkFQKpNWTMdxIoRbpHzy-W_fntVxalx1MFNd3xO27KF3pkjYvCEhCd--QtZ-S087Sw5Ja_2O3MGhTr2mPMgeY8M3aP1z4gKPjmyfxolg",
								"width" : 4032
						}
					],
					"rating" : 4.3
				}
			],
			"debug_log" : {
				"line" : []
			},
			"status" : "OK"
		}
	`

	var place Entity.Place

	var res interface{}

	if err := json.Unmarshal([]byte(myJson), &res); err != nil {
		log.Fatal(err)
	}

	address := res.(map[string]interface{})["candidates"].([]interface{})[0].(map[string]interface{})["formatted_address"].(string)
}

type TextSearch struct{}

func (t TextSearch) formatquery(queries []string) string {
	formattedQuery := queries[0]
	for _, query := range queries[1:] {
		formattedQuery = formattedQuery + "+" + query
	}

	return formattedQuery
}

func (t TextSearch) baseURL() string {
	return "https://maps.googleapis.com/maps/api/place/textsearch/json?query="
}

func (t TextSearch) options() string {
	return "&language=ja"
}

func client() {
	client := APIClient{
		TextSearch{},
	}
	client.get([]string{"aaa", "bbb"})
}

func main() {

}
