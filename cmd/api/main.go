package main

import (
	"backend/internal/repository"
	"backend/internal/repository/dbrepo"
	"backend/pkg"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

const port = 8080

type application struct {
	Domain string
	Dsn    string
	DB     repository.Repository
}

func main() {

	fmt.Println("Hello World")
	var app application
	//set config

	//read the command line
	flag.StringVar(&app.Dsn, "dsn", "host=localhost port=5432 user=postgres password=postgres dbname=movies sslmode=disable timezone=UTC connect_timeout=5", "PostgreSQL dsn")
	flag.Parse()

	//connect to the database
	app.Domain = "example.com"
	log.Println("Listening on port " + strconv.Itoa(port))

	connection, err := app.connect()
	if err != nil {
		log.Fatal(err)
	}
	app.DB = &dbrepo.PostgresDBRepo{DB: connection}
	defer pkg.GetDB().Close()
	//start the webserver
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())

	if err != nil {
		panic(err)
	}
}
