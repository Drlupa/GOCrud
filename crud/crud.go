package crud

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type User struct {
	ID    int
	Name  string
	Email string
}

const (
	dbDriver = "mysql"
	dbUser   = "gocrudapi"
	dbPass   = "admin"
	dbName   = "gocrudapi"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	var user User

	json.NewDecoder(r.Body).Decode(&user)

	err = CreateUser(db, user.Name, user.Email)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "User created successfully")
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	variables := mux.Vars(r)
	idStr := variables["id"]

	userID, err := strconv.Atoi(idStr)
	if err != nil {
		panic(err.Error())
	}

	user, err := GetUserById(db, userID)
	if err != nil {
		http.Error(w, "Failed to get user or user not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(user)
}

func CreateUser(db *sql.DB, name, email string) error {
	query := "INSERT INTO users (name, email) VALUES (?, ?)"
	_, err := db.Exec(query, name, email)
	if err != nil {
		return err
	}
	return nil
}

func GetUserById(db *sql.DB, id int) (*User, error) {
	query := "SELECT * FROM users WHERE id=?"
	row := db.QueryRow(query, id)
	user := &User{}
	row.Scan(&user.ID, &user.Name, &user.Email)
	return user, nil
}
