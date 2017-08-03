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

func PostUser(u User) (int, error) {
	db, err := getConnection()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	st, err := db.Prepare("call SIS_user_insert();")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer st.Close()

	rows, err := st.Execute()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer rows.Close()

	
	return 1, err
}

func PutUser() (int, error) {
	db, err := getConnection()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	st, err := db.Prepare("call SIS_user_update();")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer st.Close()

	rows, err := st.Execute()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer rows.Close()

	
	return 1, err
}

func DeleteUser(id int) (int, error) {
	db, err := getConnection()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	st, err := db.Prepare("call SIS_user_delete();")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer st.Close()

	rows, err := st.Execute(iid)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer rows.Close()

	
	return 1, err
}

func GetUser() (User, error) {
	db, err := getConnection()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	st, err := db.Prepare("call SIS_user_detail();")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer st.Close()

	row := st.QueryRow()
	var user User
	row.Scan(&user.Name, &user.Old)

	return user, err
}
