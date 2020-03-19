package main

import (
	"fmt"
	"net/http"

	userscontroller "github.com/ajangi/gAuthService/app/controllers/users"
	verificationcodecontroller "github.com/ajangi/gAuthService/app/controllers/verification"
	"github.com/ajangi/gAuthService/app/types/responses"
	"github.com/labstack/echo"
)
func main() {
	echo.NotFoundHandler = func(c echo.Context) error {
		errorResponse := responses.GetPageNotFoundError()
		return c.JSON(http.StatusNotFound, errorResponse)
	}
	e := echo.New()
	e.HTTPErrorHandler = customHTTPErrorHandler
	v1ApiGroup := e.Group("/api/v1")
	usersGroup := v1ApiGroup.Group("/users")
	usersGroup.POST("/getCode", verificationcodecontroller.GetVerificationCode)
	usersGroup.POST("/register", userscontroller.RegisterUser)
	e.Logger.Fatal(e.Start(":8020"))
}

func customHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		fmt.Println("error",he.Error())
	}
	var ErrorEn, ErrorFa = "", ""
	if code == 500 {
		ErrorEn = "Internal Server Error!"
		ErrorFa = "مشکلی در اتصال به سرور رخ داده است."
	} else {
		ErrorEn = err.Error()
		ErrorFa = err.Error()
	}
	errorResponse := responses.GetCustomHTTPError(ErrorEn, ErrorFa, code)
	if err = c.JSON(code, errorResponse); err != nil {
		c.Logger().Error(err)
	}
	c.Logger().Error(err)
}
