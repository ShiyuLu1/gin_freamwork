package models

type BaseResponse struct {
	Code    int                    `json:"errno"`
	Message string                 `json:"errmsg"`
	Data    map[string]interface{} `json:"data"`
}
