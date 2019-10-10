package googlemap

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	Entity "github.com/yumuranaoki/echo-practice/entity"
)

type APIClient struct {
	Strategy
}

type Strategy interface {
	formatquery([]string) string
	baseURL() string
	options() string
}

func (client APIClient) Get(keywords []string) Entity.Place {
	apiKey := os.Getenv("APIKEY")
	url := client.baseURL() + client.formatquery([]string{"東京", "観光"}) + client.options() + "&key=" + apiKey
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

	address := res.(map[string]interface{})["results"].([]interface{})[0].(map[string]interface{})["formatted_address"].(string)
	name := res.(map[string]interface{})["results"].([]interface{})[0].(map[string]interface{})["name"].(string)
	placeID := res.(map[string]interface{})["results"].([]interface{})[0].(map[string]interface{})["place_id"].(string)
	photoReference := res.(map[string]interface{})["results"].([]interface{})[0].(map[string]interface{})["photos"].([]interface{})[0].(map[string]interface{})["photo_reference"].(string)

	place := Entity.Place{
		Address:        address,
		Name:           name,
		PlaceID:        placeID,
		PhotoReference: photoReference,
	}

	return place
}
