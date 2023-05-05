package response

// APIResponse ...
type APIResponse struct {
	Status        string      `json:"status"`
	StatusCode    int         `json:"statusCode"`
	StatusMessage string      `json:"statusMessage"`
	Data          interface{} `json:"data,omitempty"`
}
