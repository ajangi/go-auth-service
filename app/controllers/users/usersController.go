package userscontroller

import (
	"fmt"
	"net/http"

	"github.com/ajangi/gAuthService/app/utils/types"
	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"
)

var validate *validator.Validate

// GetVerificationCode is the main function for getting verification code
func GetVerificationCode(c echo.Context) (err error) {
	validate = validator.New()
	request := new(types.GetCodeRequest)
	if err = c.Bind(request); err != nil {
		return
	}
	if err = validate.Struct(c.Bind(request)); err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return err
		}
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err.Namespace())
			fmt.Println(err.Field())
			fmt.Println(err.StructNamespace())
			fmt.Println(err.StructField())
			fmt.Println(err.Tag())
			fmt.Println(err.ActualTag())
			fmt.Println(err.Kind())
			fmt.Println(err.Type())
			fmt.Println(err.Value())
			fmt.Println(err.Param())
			return c.JSON(http.StatusBadRequest, err.Value())
		}
	}
	return c.JSON(http.StatusOK, request)
}
