package handler

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/yumuranaoki/date/entity/googlemap"
	"github.com/yumuranaoki/date/entity/googlemap/strategy"
)

func RestaurantHandler(c echo.Context) (err error) {
	placeIDStrategy := strategy.PlaceID{}

	IDclient := googlemap.APIClient{
		Strategy: placeIDStrategy,
	}

	placeID := c.QueryParam("place_id")
	locality := IDclient.Get([]string{placeID})[0].LongName

	textSearchStrategy := strategy.TextSearch{}
	textSearchClient := googlemap.APIClient{
		Strategy: textSearchStrategy,
	}

	mapInformations := textSearchClient.Get([]string{locality, "レストラン"})

	return c.JSON(http.StatusOK, mapInformations)
}
