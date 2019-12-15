package main

import (
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/yumuranaoki/date/handler"
)

func newRouter() *echo.Echo {
	e := echo.New()

	e.GET("/leisures", handler.LeisureHandler)
	e.GET("/restaurants", handler.RestaurantHandler)
	e.GET("/summary", handler.SummaryHandler)

	return e
}

func main() {
	router := newRouter()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router.Use(middleware.Logger())

	router.Start(":" + port)
}
