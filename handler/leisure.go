package handler

import (
	"net/http"

	"github.com/labstack/echo"
	Entity "github.com/yumuranaoki/echo-practice/entity"
)

func LeisureHandler(c echo.Context) (err error) {
	place := new(Entity.Place)

	if err = c.Bind(place); err != nil {
		return
	}

	return c.JSON(http.StatusOK, place)
}
