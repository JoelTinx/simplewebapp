package handlers

import "github.com/JoelTinx/sismplewebapp/models"

const (

)

func GetallUser() {
	users, err := models.GetallUser()
	if err != nil {
		panic(err)
	}
}

func PostUser() {

}

func UpdateUser() {

}

func DeleteUser(cCodUsuSis string) {

}

func ChangePassword() {

}

func ResetPassword()  {

}
