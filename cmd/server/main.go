package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"

	"github.com/Pleopleq/group-chat/internal/api"
)

func main() {
	server := &http.Server{
		Addr: ":8080",
	}

	db, err := sql.Open("sqlite3", "../../db/app.db")
	if err != nil {
		log.Fatal(fmt.Errorf("got an error creating a connection to the DB, %v", err))
	}
	defer db.Close()

	http.HandleFunc("/api/v1/login", api.Login)
	http.HandleFunc("/api/v1/register", api.Register)

	log.Fatal(server.ListenAndServe())
}
