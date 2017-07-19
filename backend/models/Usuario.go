package models

type User struct {
	Name string `json:"name"`
	Old int `json:"old"`
}

func GetAllUser() ([]User, error){
	conn, err := getConnection()
	if err != nil {
		panic(err)
	}
}

func PostUser() (int, error) {

}

func PutUser() (int, error) {

}

func DeleteUser() (int, error) {

}

func GetUser() (User, error) {

}
