package handler

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/yumuranaoki/date/entity/googlemap"
	"github.com/yumuranaoki/date/entity/googlemap/strategy"
)

func LeisureHandler(c echo.Context) (err error) {
	textSearch := strategy.TextSearch{}

	client := googlemap.APIClient{
		Strategy: textSearch,
	}

	place := c.QueryParam("place")

	mapInformations := client.Get([]string{place, "観光"})

	return c.JSON(http.StatusOK, mapInformations)
}
