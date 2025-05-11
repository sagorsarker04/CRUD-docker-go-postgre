package main

import (
	"fmt"
	"log"
	"net/http"

	"todo/config"
	"todo/db"
	"todo/handlers"

	"github.com/gorilla/mux"
)

func main() {
	cfg := config.LoadConfig()
	db.InitDB(cfg)
	r := mux.NewRouter()
	r.HandleFunc("/users", handlers.CreateUser).Methods("POST")
	r.HandleFunc("/table", handlers.CreateTable).Methods("GET")
	r.HandleFunc("/users", handlers.GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", handlers.GetUserByID).Methods("GET")
	r.HandleFunc("/users/{id}", handlers.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", handlers.DeleteUser).Methods("DELETE")

	fmt.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
