package models

//定义User类型结构
type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

