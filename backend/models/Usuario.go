package models

type User struct {
	Name string `json:"name"`
	Old int `json:"old"`
}

func GetAllUser() ([]User, error){
	
}