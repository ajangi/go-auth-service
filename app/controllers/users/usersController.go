package userscontroller

import (
	"net/http"
	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"
	User "github.com/ajangi/gAuthService/app/models"
	"github.com/ajangi/gAuthService/app/types/requests"
	"github.com/ajangi/gAuthService/app/types/responses"
	"github.com/ajangi/gAuthService/app/services/validation"
)

var validate *validator.Validate

// RegisterUser is the method that gets user register request and registers user or returns error response
func RegisterUser(c echo.Context) (err error)  {
	validate = validator.New()
	validate.RegisterValidation("uniqueInDB", validation.UniqueInDB)
	request := new(requests.RegisterUserRequest)
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
	user,err := User.CreateUser(request)
	if err != nil {
		return
	}
	return c.JSON(http.StatusOK, user)
}