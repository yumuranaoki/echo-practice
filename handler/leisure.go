package handler

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/yumuranaoki/echo-practice/entity/googlemap"
)

func LeisureHandler(c echo.Context) (err error) {
	// place := new(Entity.Place)

	// if err = c.Bind(place); err != nil {
	// 	return
	// }

	textSearch := googlemap.TextSearch{}

	client := googlemap.APIClient{
		Strategy: textSearch,
	}

	place := client.Get([]string{"tokyo"})
	// client.get

	return c.JSON(http.StatusOK, place)
}
