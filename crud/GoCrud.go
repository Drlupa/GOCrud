package crud

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
    _ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID int
	Name string
	Email string
}

const (
	dbDriver  = "mysql"
	dbUser = "gocrudapi"
	dbPass = "admin"
	dbName = "gocrudapi"
)

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
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

func CreateUser(db *sql.DB, name, email string) error {
    query := "INSERT INTO users (name, email) VALUES (?, ?)"
    _, err := db.Exec(query, name, email)
    if err != nil {
        return err
    }
    return nil
}