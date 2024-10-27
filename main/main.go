package main

import (
	"log"
	"net/http"

	"github.com/Drlupa/GOCrud/crud"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/user", crud.CreateUserHandler).Methods("POST")
	r.HandleFunc("/user/{id}", crud.GetUserHandler).Methods("GET")
	r.HandleFunc("/user/{id}", crud.UpdateUserHandler).Methods("PUT")
	r.HandleFunc("/user/{id}", crud.DeleteUserHandler).Methods("DELETE")

	log.Println("Server listening on :8090")
	log.Fatal(http.ListenAndServe(":8090", r))
}
