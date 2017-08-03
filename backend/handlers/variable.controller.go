package handlers

func GetVariable(cNomVarSis string) (Variable, error){
	db, err := getConnection()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	st, err := db.Prepare("SELECT SIS_variable_value(?);")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer st.Close()

	row := st.QueryRow(cNomVarSis)
	var v Variable
	row.Scan(&user.Name, &user.Old)

	return v, err
}

func PutVariable() {
	
}