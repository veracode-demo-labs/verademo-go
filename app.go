package main

import (
	"database/sql"
	"log"
	"net/http"
	router "verademo-go/src-app/routes"
	db "verademo-go/src-app/shared/db"
)

func main() {
	var database *sql.DB
	database, _ = db.InitDB()
	log.Print("test")

	log.Fatal(http.ListenAndServe(":8080", router.Routes()))
	database.Close()
}
