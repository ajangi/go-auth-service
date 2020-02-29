package main

import (
	"net/http"

	"github.com/ajangi/gAuthService/app/utils/responses"
	"github.com/labstack/echo"
)

func main() {
	echo.NotFoundHandler = func(c echo.Context) error {
		errorResponse := responses.GetPageNotFoundError()
		return c.JSON(http.StatusNotFound, errorResponse)
	}
	e := echo.New()
	usersGroup := e.Group("/users")
	usersGroup.POST("", nil)
	e.Logger.Fatal(e.Start(":8020"))
}
