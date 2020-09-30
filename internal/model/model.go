package model

type User struct {
	Id int `json:"id" gorm:"primary_key"`
	Account string `json:"account"`
	Password string `json:"password"`
}

func NewUser() *User {
	return &User{
		Id:       0,
		Account:  "",
		Password: "",
	}
}