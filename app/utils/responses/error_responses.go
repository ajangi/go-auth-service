package responses

import (
	"github.com/ajangi/gAuthService/app/utils/types"
	"net/http"
)

// GetPageNotFoundError function is to export page not found error response
func GetPageNotFoundError() types.ApiResponse {
	messages := map[string]types.MessageItem{
		"general": types.MessageItem{EN:[]string{"Requested Page Not Found!"},FA:[]string{"صفحه مورد نظر بافت نشد."}},
	}
	emptyData := types.ResponseData{}
	return types.ApiResponse{
		Result: "ERROR",
		StatusCode: http.StatusNotFound,
		Messages: messages,
		Data: emptyData,
	}
}

// GetCustomHTTPError function is to override echo custom errors
func GetCustomHTTPError(ErrorStringEn string, ErrorStringFa string, StatusCode int) types.ApiResponse {
	messages := map[string]types.MessageItem{
		"general": types.MessageItem{EN:[]string{ErrorStringEn},FA:[]string{ErrorStringFa}},
	}
	emptyData := types.ResponseData{}
	return types.ApiResponse{
		Result: "ERROR",
		StatusCode: StatusCode,
		Messages: messages,
		Data: emptyData,
	}
}