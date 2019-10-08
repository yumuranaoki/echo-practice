package main

import (
	"github.com/labstack/echo"

	"github.com/yumuranaoki/echo-practice/handler"
)

func newRouter() *echo.Echo {
	e := echo.New()

	e.GET("/leisures", handler.LeisureHandler)

	return e
}

func main() {
	router := newRouter()

	router.Start(":8000")
}
