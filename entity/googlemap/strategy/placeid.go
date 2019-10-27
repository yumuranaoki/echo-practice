package strategy

import (
	Entity "github.com/yumuranaoki/date/entity"
)

type PlaceID struct{}

func (p PlaceID) Formatquery(keywords []string) string {
	return keywords[0]
}

func (p PlaceID) BaseURL() string {
	return "https://maps.googleapis.com/maps/api/place/details/json?placeid="
}

func (p PlaceID) Options() string {
	return "&language=ja"
}

func (p PlaceID) Parse(res interface{}) []Entity.Place {
	var places []Entity.Place = []Entity.Place{}
	var result interface{} = res.(map[string]interface{})["result"]

	var addressComponents interface{} = result.(map[string]interface{})["address_components"]
	switch addressComponents.(type) {
	case []interface{}:
		for _, component := range addressComponents.([]interface{}) {
			longName := component.(map[string]interface{})["long_name"].(string)
			types := component.(map[string]interface{})["types"].([]string)

			if includeLocality(types) {
				place := Entity.Place{
					Name:           longName,
					Address:        "",
					PlaceID:        "",
					PhotoReference: "",
				}
				places = []Entity.Place{place}

				break
			}
		}
	}

	return places
}

func includeLocality(arr []string) bool {
	for _, el := range arr {
		if el == "locality" {
			return true
		}
	}

	return false
}
