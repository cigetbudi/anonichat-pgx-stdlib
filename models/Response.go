package models

type Response struct {
	StatusCode  string      `json:"status_code"`
	Description string      `json:"description"`
	Data        interface{} `json:"data"`
}

type TokenResp struct {
	Token string `json:"token"`
}
