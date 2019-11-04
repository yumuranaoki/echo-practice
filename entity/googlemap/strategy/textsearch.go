package strategy

import (
	Entity "github.com/yumuranaoki/date/entity"
)

type TextSearch struct{}

func (t TextSearch) Formatquery(queries []string) string {
	formattedQuery := queries[0]
	for _, query := range queries[1:] {
		formattedQuery = formattedQuery + "+" + query
	}

	return formattedQuery
}

func (t TextSearch) BaseURL() string {
	return "https://maps.googleapis.com/maps/api/place/textsearch/json?query="
}

func (t TextSearch) Options() string {
	return "&language=ja"
}

func (t TextSearch) Parse(res interface{}) []Entity.MapInformation {
	var mapInformations []Entity.MapInformation
	var results interface{} = res.(map[string]interface{})["results"]

	switch results.(type) {
	case []interface{}:
		for _, result := range results.([]interface{}) {
			address := result.(map[string]interface{})["formatted_address"].(string)
			name := result.(map[string]interface{})["name"].(string)
			placeID := result.(map[string]interface{})["place_id"].(string)
			photoReference := result.(map[string]interface{})["photos"].([]interface{})[0].(map[string]interface{})["photo_reference"].(string)

			mapInformation := Entity.MapInformation{
				Address:        address,
				Name:           name,
				PlaceID:        placeID,
				PhotoReference: photoReference,
			}

			mapInformations = append(mapInformations, mapInformation)
		}
	default:
		mapInformations = []Entity.MapInformation{}
	}

	return mapInformations
}
