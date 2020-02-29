package main

import (
	"net/http"

	userscontroller "github.com/ajangi/gAuthService/app/controllers/users"
	"github.com/ajangi/gAuthService/app/utils/responses"
	"github.com/labstack/echo"
)

func main() {
	echo.NotFoundHandler = func(c echo.Context) error {
		errorResponse := responses.GetPageNotFoundError()
		return c.JSON(http.StatusNotFound, errorResponse)
	}
	e := echo.New()
	v1ApiGroup := e.Group("/api/v1")
	usersGroup := v1ApiGroup.Group("/users")
	usersGroup.POST("/getCode", userscontroller.GetVerificationCode)
	e.Logger.Fatal(e.Start(":8020"))
}
