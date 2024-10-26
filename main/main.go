package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
    _ "github.com/go-sql-driver/mysql"
    "github.com/gorilla/mux"
	"github.com/Drlupa/GOCrud"
)

const (
	dbDriver  = "mysql"
	dbUser = "gocrudapi"
	dbPass = "admin"
	dbName = "gocrudapi"
)

func main() {
	r:= mux.NewRouter()

	r.HandleFunc("/user", createUserHandler).Methods("POST")
    // r.HandleFunc("/user/{id}", getUserHandler).Methods("GET")
    // r.HandleFunc("/user/{id}", updateUserHandler).Methods("PUT")
    // r.HandleFunc("/user/{id}", deleteUserHandler).Methods("DELETE")

	log.Println("Server listening on :8090")
    log.Fatal(http.ListenAndServe(":8090", r))
	}