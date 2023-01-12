package model

type WebResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type PaginateResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Next    string      `json:"next"`
	Data    interface{} `json:"data"`
}

type TokenResponse struct {
	Type   string      `json:"type"`
	Alg    string      `json:"alg"`
	Token  string      `json:"token"`
	Claims interface{} `json:"claims"`
}
