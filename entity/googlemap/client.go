package googlemap

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	Entity "github.com/yumuranaoki/date/entity"
)

type APIClient struct {
	Strategy
}

type Strategy interface {
	Formatquery([]string) string
	BaseURL() string
	Options() string
}

func (client APIClient) Get(keywords []string) []Entity.Place {
	apiKey := os.Getenv("APIKEY")
	url := client.BaseURL() + client.Formatquery(keywords) + client.Options() + "&key=" + apiKey
	resp, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var res interface{}

	if err := json.Unmarshal(body, &res); err != nil {
		log.Fatal(err)
	}

	var places []Entity.Place
	var results interface{} = res.(map[string]interface{})["results"]

	switch results.(type) {
	case []interface{}:
		for _, result := range results.([]interface{}) {
			address := result.(map[string]interface{})["formatted_address"].(string)
			name := result.(map[string]interface{})["name"].(string)
			placeID := result.(map[string]interface{})["place_id"].(string)
			photoReference := result.(map[string]interface{})["photos"].([]interface{})[0].(map[string]interface{})["photo_reference"].(string)

			place := Entity.Place{
				Address:        address,
				Name:           name,
				PlaceID:        placeID,
				PhotoReference: photoReference,
			}

			places = append(places, place)
		}
	default:
	}

	return places
}
