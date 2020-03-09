package responses

import (
	"github.com/ajangi/gAuthService/app/types"
	"net/http"
)

// GetPageNotFoundError function is to export page not found error response
func GetPageNotFoundError() APIResponse {
	messages := map[string]MessageItem{
		"general": MessageItem{EN:[]string{"Requested Page Not Found!"},FA:[]string{"صفحه مورد نظر بافت نشد."}},
	}
	emptyData := ResponseData{}
	return APIResponse{
		Result: "ERROR",
		StatusCode: http.StatusNotFound,
		Messages: messages,
		Data: emptyData,
	}
}

// GetCustomHTTPError function is to override echo custom errors
func GetCustomHTTPError(ErrorStringEn string, ErrorStringFa string, StatusCode int) APIResponse {
	messages := map[string]MessageItem{
		"general": MessageItem{EN:[]string{ErrorStringEn},FA:[]string{ErrorStringFa}},
	}
	emptyData := ResponseData{}
	return APIResponse{
		Result: "ERROR",
		StatusCode: StatusCode,
		Messages: messages,
		Data: emptyData,
	}
}


// GetValidationErrorResponse function is to generate validation error Response 
func GetValidationErrorResponse(errorsList types.ErrorsList) APIResponse {
	messages := map[string]MessageItem{}
		for i,v := range errorsList{
			messages[i] = MessageItem{FA:[]string{v.Fa},EN:[]string{v.En}}
		}
		errorResponse := APIResponse{}
		errorResponse.Data = ResponseData{}
		errorResponse.Result = "ERROR"
		errorResponse.StatusCode = http.StatusBadRequest
		errorResponse.Messages = messages
		return errorResponse
}