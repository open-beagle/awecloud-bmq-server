package model

type APIResponse struct {
	Success int         `json:"success"`
	ErrMsg  string      `json:"errMsg,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type ArrayResponse struct {
	Success int         `json:"success"`
	ErrMsg  string      `json:"errMsg,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Total   int         `json:"total"`
}
