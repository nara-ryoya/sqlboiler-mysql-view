package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nara-ryoya/sqlboiler-mysql-view/db"
	"github.com/nara-ryoya/sqlboiler-mysql-view/handler"
)

func main() {
	if err := db.Init("user", "mysql", "localhost", 3306); err != nil {
		log.Fatalf("db init: %v", err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/users", handler.Users).Methods("GET")

	log.Println("LISTEN :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
