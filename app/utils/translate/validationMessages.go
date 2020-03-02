package translate

import(
	"gopkg.in/go-playground/validator.v9"
)

// GetValidationMessage is the method that gets a message using validation error
func GetValidationMessage(err validator.FieldError) (VFMessage ValidationFieldMessage) {
	validationField := Fields[err.Field()]
	var message = ValidationFieldMessage{}
	switch err.Tag(){
		case "required":
			message = ValidationFieldMessage{Fa:validationField.Fa+" ارسال نشده است",En:validationField.En+ " is required"}
		case "len":
			message = ValidationFieldMessage{Fa:validationField.Fa+" باید دارای "+err.Param()+" کاراکتر باشد",En:validationField.En+ " should have "+err.Param()+" characters"}
	}
	return message
}