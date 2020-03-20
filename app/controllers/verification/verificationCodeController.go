package verificationcodecontroller

import (
	"github.com/ajangi/gAuthService/app/utils/db"
	"net/http"

	"github.com/ajangi/gAuthService/app/models"
	"github.com/ajangi/gAuthService/app/services/validation"
	"github.com/ajangi/gAuthService/app/types/requests"
	"github.com/ajangi/gAuthService/app/types/responses"
	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"
)

var validate *validator.Validate

// GetVerificationCode is the main function for getting verification code
func GetVerificationCode(c echo.Context) (err error) {
	validate = validator.New()
	request := new(requests.GetCodeRequest)
	var response responses.APIResponse
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
	database,err := db.GetDB()
	if err != nil {
		response = responses.GetInternalServiceError()
		return c.JSON(response.StatusCode,response)
	}
	var code models.VerificationCode
	var user models.User
	if err = database.Table(user.TableName()).Where("phone_number=?",request.PhoneNumber).First(&user).Error; err != nil{
		if err.Error() == "record not found" {
			response = responses.GetUserNotFoundError()
			return c.JSON(response.StatusCode,response)
		}
	}
	if err = database.Table(code.TableName()).Where("user_id=?",user.ID).Where("status=?","fresh").First(&code).Error; err != nil{
		if err.Error() == "record not found" {
			if err = models.CreateCode(user); err != nil{
				response = responses.GetInternalServiceError()
				return c.JSON(response.StatusCode,response)
			}
		}
	}
	response = responses.GetSimpleSuccessResponse()
	return c.JSON(response.StatusCode,response)
}
