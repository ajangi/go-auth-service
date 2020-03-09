package verificationcodecontroller

import (
	"net/http"

	"github.com/ajangi/gAuthService/app/types/requests"
	"github.com/ajangi/gAuthService/app/types/responses"
	"github.com/ajangi/gAuthService/app/services/validation"
	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"
)

var validate *validator.Validate

// GetVerificationCode is the main function for getting verification code
func GetVerificationCode(c echo.Context) (err error) {
	validate = validator.New()
	request := new(requests.GetCodeRequest)
	if err = c.Bind(request); err != nil {
		return
	}
	if err = validate.Struct(request); err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return
		}
		errorsList := validation.GetValidationErrors(err)
		errorResponse := responses.GetValidationErrorResponse(errorsList)
		return c.JSON(http.StatusBadRequest, errorResponse)
	}
	return c.JSON(http.StatusOK, request)
}
