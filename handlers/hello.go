package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func HelloHandler(c echo.Context) error {
	user := User{
		Name:  "naoki",
		Email: "naoki.yumura@gmail.com",
	}
	return c.JSON(http.StatusOK, user)
}
