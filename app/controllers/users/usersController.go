package userscontroller

import (
	"net/http"

	"github.com/ajangi/gAuthService/app/utils/types"
	"github.com/labstack/echo"
)

// GetVerificationCode is the main function for getting verification code
func GetVerificationCode(c echo.Context) (err error) {
	request := new(types.GetCodeRequest)
	if err = c.Bind(request); err != nil {
		return
	}
	return c.JSON(http.StatusOK, request)
}
