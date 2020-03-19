package translate

import (
	"fmt"
	"strings"

	"gopkg.in/go-playground/validator.v9"
)

// GetValidationMessage is the method that gets a message using validation error
func GetValidationMessage(err validator.FieldError) (VFMessage ValidationFieldMessage) {
	validationField := Fields[err.Field()]
	var message = ValidationFieldMessage{}
	switch err.Tag() {
	case "required":
		message = ValidationFieldMessage{Fa: validationField.Fa + " ارسال نشده است", En: validationField.En + " is required"}
	case "len":
		message = ValidationFieldMessage{Fa: validationField.Fa + " باید دارای " + err.Param() + " کاراکتر باشد", En: validationField.En + " should have " + err.Param() + " characters"}
	case "required_without":
		params := strings.Split(err.Param(), " ")
		paramsTranslations := map[int]ValidationField{}
		for i, param := range params {
			paramsTranslations[i] = Fields[param]
		}
		FaString := validationField.Fa + " در صورت عدم ارسال "
		EnString := validationField.En + " is required if "
		for i, translated := range paramsTranslations {
			FaString = FaString + "(" + translated.Fa + ")"
			EnString = EnString + "(" + translated.En + ")"
			if i < (len(paramsTranslations) - 1) {
				FaString = FaString + " و "
				EnString = EnString + ", "
			}
		}
		FaString = FaString + " الزامی است"
		EnString = EnString + " is not present"
		fmt.Println(paramsTranslations)
		message = ValidationFieldMessage{Fa: FaString, En: EnString}
	case "uniqueInDB":
		FaString := validationField.Fa + " قبلا استفاده شده است"
		EnString := validationField.En + " is taken before"
		message = ValidationFieldMessage{Fa: FaString, En: EnString}
	}
	return message
}
