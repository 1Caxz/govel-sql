package model

type PaginateResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Next    string      `json:"next"`
	Data    interface{} `json:"data"`
}
