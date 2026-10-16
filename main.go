package main

import (
	"log"

	"github.com/Vijay6922/Go-DBconnection/cmd/api"
	"github.com/Vijay6922/Go-DBconnection/sqlMethods"
)

func main() {
	db, err := sqlMethods.OpenDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	api.API(db)
}
