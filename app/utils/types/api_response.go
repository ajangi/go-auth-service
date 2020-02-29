package types

type ApiResponse struct {
	Result string `json:"result"`
	StatusCode int `json:"status_code"`
	Messages map[string]MessageItem `json:"messages"`
	Data ResponseData `json:"data"`
}

type MessageItem struct {
	EN []string `json:"en"`
	FA []string `json:"fa"`
}

type ResponseData []interface{}