package googlemap

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	Entity "github.com/yumuranaoki/date/entity"
)

// APIClient is client for googlemap api
type APIClient struct {
	Strategy
}

// Strategy specify how to generate URL
type Strategy interface {
	Formatquery([]string) string
	BaseURL() string
	Options() string
	Parse(interface{}) []Entity.Place
}

// Get returns place information
func (client APIClient) Get(keywords []string) []Entity.Place {
	url := client.generateURL(keywords)
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

	places := client.Parse(res)

	return places
}

func (client APIClient) generateURL(keywords []string) string {
	apiKey := os.Getenv("APIKEY")

	return client.BaseURL() + client.Formatquery(keywords) + client.Options() + "&key=" + apiKey
}
