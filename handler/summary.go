package handler

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/yumuranaoki/date/entity/googlemap"
	"github.com/yumuranaoki/date/entity/googlemap/strategy"
)

func SummaryHandler(c echo.Context) (err error) {
	DirectionStrategy := strategy.Direction{}

	DirectionClient := googlemap.APIClient{
		Strategy: DirectionStrategy,
	}

	originPlaceID := c.QueryParam("origin_place_id")
	destinationPlaceID := c.QueryParam("destination_place_id")
	mapInformation := DirectionClient.Get([]string{originPlaceID, destinationPlaceID})[0]

	return c.JSON(http.StatusOK, mapInformation)
}
