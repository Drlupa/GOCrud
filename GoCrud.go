package GoCrud


type USER struct {
	ID int
	Name string
	Email string
}


func createUserHandler(w http.ResponseWriter, r *http.Request) {
	db, err = sql.open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	var user User

	json.NewDecoder(r.Body).Decode(&user)

	CreateUser(db, user.Name, user.Email)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	   }
   
	   w.WriteHeader(http.StatusCreated)
	   fmt.Fprintln(w, "User created successfully")
}