package main

import (
	"github.com/labstack/echo"
)

func main() {
	router := NewRouter()

	router.Start(":8000")
}

func NewRouter() *echo.Echo {
	e := echo.New()

	e.GET("/hello", handler)

	return e
}
