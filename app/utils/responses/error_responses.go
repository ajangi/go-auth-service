package responses

import (
	"../types"
	"net/http"
)

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