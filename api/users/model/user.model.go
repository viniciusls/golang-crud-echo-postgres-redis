package model

type User struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
	Age      int64  `json:"age"`
}
