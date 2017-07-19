package models

type User struct {
	Name string `json:"name"`
	Old int `json:"old"`
}

func GetAllUser() ([]User, error) {
	db, err := getConnection()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	st, err := db.Prepare("call SIS_user_list();")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer st.Close()

	rows, err := st.Query()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer rows.Close()

	users := make([]User, 0)
	for rows.Next() {
		var u User

		err = rows.Scan(&u.Name, &u.Old)
		users = append(users, u)
	}
	return users, err
}

func PostUser() (int, error) {

}

func PutUser() (int, error) {

}

func DeleteUser() (int, error) {

}

func GetUser() (User, error) {

}
