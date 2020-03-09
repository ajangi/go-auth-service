package validation

import (
	"fmt"

	"github.com/ajangi/gAuthService/app/types"
	"github.com/ajangi/gAuthService/app/utils/translate"
	"gopkg.in/go-playground/validator.v9"
)

// GetValidationErrors is a function to get validation errors
func GetValidationErrors(err error) types.ErrorsList {
	els := types.ErrorsList{}
	for _, err := range err.(validator.ValidationErrors) {
		els[err.Field()] = translate.GetValidationMessage(err)
		fmt.Println("namespace: ",err.Namespace())
		fmt.Println("field: ",err.Field())
		fmt.Println("struct namespace: ",err.StructNamespace())
		fmt.Println("struct field: ",err.StructField())
		fmt.Println("tag: ",err.Tag())
		fmt.Println("actual tag: ",err.ActualTag())
		fmt.Println("kind: ",err.Kind())
		fmt.Println("type: ",err.Type())
		fmt.Println("value: ", err.Value())
		fmt.Println("param: ",err.Param())
		fmt.Println("----------------------")
	}
	return els
}


// ValidateGetCodeRequest is a function that validates get code request
func ValidateGetCodeRequest(){

}

