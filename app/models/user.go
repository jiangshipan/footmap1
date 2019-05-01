package models

type User struct {
	Uid int
	Username string `json:"username"`
	Password string
	Nickname string `json:"nickname"`
	Status int
	Salt string
	Manager int
}
type ResponseUser struct {
	Nickname string `json:"nickname"`
	Manager int `json:"manager"`
}