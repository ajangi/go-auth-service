package responses

// APIResponse is the main response struct that should be returned in api
type APIResponse struct {
	Result string `json:"result"`
	StatusCode int `json:"status_code"`
	Messages map[string]MessageItem `json:"messages"`
	Data ResponseData `json:"data"`
}

// MessageItem is the message item that includes en and fa array of sttrings
type MessageItem struct {
	EN []string `json:"en"`
	FA []string `json:"fa"`
}

// ResponseData is the data that should be returned in response
type ResponseData []interface{}