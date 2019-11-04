package strategy

import (
	"log"
	"strconv"

	Entity "github.com/yumuranaoki/date/entity"
)

type Direction struct{}

func (d Direction) Formatquery(queries []string) string {
	if len(queries) != 2 {
		log.Fatalf("expected 2 args, but got %s", strconv.Itoa(len(queries)))
	}

	return "origin=place_id:" + queries[0] + "&destination=place_id:" + queries[1]
}

func (d Direction) BaseURL() string {
	return "https://maps.googleapis.com/maps/api/directions/json?"
}

func (d Direction) Options() string {
	return "&language=ja"
}

func (d Direction) Parse(res interface{}) []Entity.MapInformation {
	routes := res.(map[string]interface{})["routes"].([]interface{})
	var duration string

	for _, route := range routes {
		legs := route.(map[string]interface{})["legs"].([]interface{})
		duration = legs[0].(map[string]interface{})["duration"].(map[string]interface{})["text"].(string)
	}

	mapInformations := Entity.MapInformation{
		Duration: duration,
	}

	return []Entity.MapInformation{mapInformations}
}
