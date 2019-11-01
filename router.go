package main

import (
	"github.com/labstack/echo"

	"github.com/yumuranaoki/date/handler"
)

func newRouter() *echo.Echo {
	e := echo.New()

	e.GET("/leisures", handler.LeisureHandler)
	e.GET("/restaurants", handler.RestaurantHandler)

	return e
}

func main() {
	router := newRouter()

	router.Start(":8000")
}
