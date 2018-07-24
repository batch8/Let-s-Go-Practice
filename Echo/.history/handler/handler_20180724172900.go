package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

func MainPage() echo.HandlerFunc {
	return func(c echo.Context) error { //c をいじって Request, Responseを色々する
		username := c.Param("username")
		return c.String(http.StatusOK, "Hello World "+username)
	}
}