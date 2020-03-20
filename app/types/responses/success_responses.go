package responses

import "net/http"

// GetSimpleSuccessResponse function is to export simple success response response
func GetSimpleSuccessResponse() APIResponse {
	messages := map[string]MessageItem{}
	emptyData := ResponseData{}
	return APIResponse{
		Result: "SUCCESS",
		StatusCode: http.StatusOK,
		Messages: messages,
		Data: emptyData,
	}
}
