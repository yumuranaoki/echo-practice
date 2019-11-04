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

func (p PlaceID) Parse(res interface{}) []Entity.MapInformation {
	var mapInformations []Entity.MapInformation = []Entity.MapInformation{}
	var result interface{} = res.(map[string]interface{})["result"]

	addressComponents := result.(map[string]interface{})["address_components"]
	switch addressComponents.(type) {
	case []interface{}:
		for _, component := range addressComponents.([]interface{}) {
			longName := component.(map[string]interface{})["long_name"].(string)
			types := component.(map[string]interface{})["types"].([]interface{})

			if includeLocality(types) {
				mapInformation := Entity.MapInformation{
					Name:           "",
					Address:        "",
					PlaceID:        "",
					PhotoReference: "",
					LongName:       longName,
				}
				mapInformations = []Entity.MapInformation{mapInformation}

				break
			}
		}
	default:
		mapInformations = []Entity.MapInformation{}
	}

	return mapInformations
}

func includeLocality(arr []interface{}) bool {
	for _, el := range arr {
		if el == "locality" {
			return true
		}
	}

	return false
}
