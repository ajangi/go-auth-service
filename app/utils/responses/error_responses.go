package responses

import (
	"github.com/ajangi/gAuthService/app/utils/types"
	"net/http"
)

// GetPageNotFoundError function is to export page not found error response
func GetPageNotFoundError() types.APIResponse {
	messages := map[string]types.MessageItem{
		"general": types.MessageItem{EN:[]string{"Requested Page Not Found!"},FA:[]string{"صفحه مورد نظر بافت نشد."}},
	}
	emptyData := types.ResponseData{}
	return types.APIResponse{
		Result: "ERROR",
		StatusCode: http.StatusNotFound,
		Messages: messages,
		Data: emptyData,
	}
}

// GetCustomHTTPError function is to override echo custom errors
func GetCustomHTTPError(ErrorStringEn string, ErrorStringFa string, StatusCode int) types.APIResponse {
	messages := map[string]types.MessageItem{
		"general": types.MessageItem{EN:[]string{ErrorStringEn},FA:[]string{ErrorStringFa}},
	}
	emptyData := types.ResponseData{}
	return types.APIResponse{
		Result: "ERROR",
		StatusCode: StatusCode,
		Messages: messages,
		Data: emptyData,
	}
}


// GetValidationErrorResponse function is to generate validation error Response 
func GetValidationErrorResponse(errorsList types.ErrorsList) types.APIResponse {
	messages := map[string]types.MessageItem{}
		for i,v := range errorsList{
			messages[i] = types.MessageItem{FA:[]string{v.Fa},EN:[]string{v.En}}
		}
		errorResponse := types.APIResponse{}
		errorResponse.Data = types.ResponseData{}
		errorResponse.Result = "ERROR"
		errorResponse.StatusCode = http.StatusBadRequest
		errorResponse.Messages = messages
		return errorResponse
}