package handler

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/yumuranaoki/date/entity/googlemap"
	"github.com/yumuranaoki/date/entity/googlemap/strategy"
)

func RestaurantHandler(c echo.Context) (err error) {
	placeIDStrategy := strategy.PlaceID{}

	client := googlemap.APIClient{
		Strategy: placeIDStrategy,
	}

	placeID := c.QueryParam("place_id")

	places := client.Get([]string{placeID})

	return c.JSON(http.StatusOK, places)
}
